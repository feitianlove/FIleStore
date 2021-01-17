package handler

import (
	"fmt"
	"github.com/feitianlove/FIleStore/models"
	"github.com/feitianlove/FIleStore/pool"
	"github.com/feitianlove/golib/common/ecode"
	"github.com/garyburd/redigo/redis"
	"os"
	"strings"

	"math"
	"net/http"
	"strconv"
	"time"
)

//TODO 分块上传和断点续传
// 1。 小文件不建议分块上传
// 2。 可以并行分块上传，并且可以无序传输
// 3。 分块上传可以极大提高传输效率
// 4。 减少传输失败后充实的流量和时间

type MultipartUploadInfo struct {
	FileHash   string
	FileSize   int
	UploadID   string
	ChunkSize  int
	ChunkCount int
}

// 初始化分块信息
func InitiateMultipartUp(w http.ResponseWriter, r http.Request) {
	//1. 判断是否已经上传过
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("Fail to parse form, err %s ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fileHash := r.Form.Get("filehash")
	userName := r.Form.Get("username")
	fileSize, _ := strconv.Atoi(r.Form.Get("filesize"))
	//2. 获得redis链接
	rConn := pool.RedisPool().Get()
	defer func() {
		_ = rConn.Close()
	}()
	//3. 生成唯一ID缓冲分块上传初始化信息
	uplodInfo := MultipartUploadInfo{
		FileHash:   fileHash,
		FileSize:   fileSize,
		UploadID:   userName + fmt.Sprintf("%x", time.Now().Unix()),
		ChunkSize:  5 * 1024 * 1024, // 5MB
		ChunkCount: int(math.Ceil(float64(fileSize) / 5 * 1024 * 1024)),
	}
	//4. 将初始化信息写入redis
	_, _ = rConn.Do("HSET", "FT"+uplodInfo.UploadID, "chunkcount", uplodInfo.ChunkCount)
	_, _ = rConn.Do("HSET", "FT"+uplodInfo.UploadID, "filehash", uplodInfo.FileHash)
	_, _ = rConn.Do("HSET", "FT"+uplodInfo.UploadID, "filesize", uplodInfo.FileSize)
	//5. 将响应初始化信息返回客户端
	resp, _ := ecode.NewResponseMsg(0, "success", uplodInfo).JOSNBytes()
	_, _ = w.Write(resp)
}

// 上传分块
func UploadPartHandler(w http.ResponseWriter, r http.Request) {
	//1. 解析用户请求参数
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("Fail to parse form, err %s ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//userName := r.Form.Get("username")
	uploadID := r.Form.Get("uploadid")
	chunkIndex := r.Form.Get("index")
	//2. 获得redis链接
	rConn := pool.RedisPool().Get()
	defer func() {
		_ = rConn.Close()
	}()
	//3. 获得文件句柄， 用与存储分块内容
	file, err := os.Create("/data/" + uploadID + "/" + chunkIndex)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		resp, _ := ecode.NewResponseMsg(-1, "upload failed", nil).JOSNBytes()
		_, _ = w.Write(resp)
	}
	defer func() {
		_ = file.Close()
	}()
	buf := make([]byte, 1024*1024*5)
	for {
		n, err := r.Body.Read(buf)
		_, _ = file.Write(buf[:n])
		if err != nil {
			break
		}
	}
	//4. 更新redis缓存状态
	_, _ = rConn.Do("HSET", "FT"+uploadID, "chkidx"+chunkIndex, 1)
	//5. 返回处理结果
	resp, _ := ecode.NewResponseMsg(0, "success", nil).JOSNBytes()
	_, _ = w.Write(resp)
}

//通知上传分块完成
func CompleteUploadHandler(w http.ResponseWriter, r http.Request) {
	// 1。解析请求参数
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("Fail to parse form, err %s ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//userName := r.Form.Get("username")
	uploadID := r.Form.Get("uploadid")
	fileHash := r.Form.Get("filehash")
	fileName := r.Form.Get("filename")
	userName := r.Form.Get("username")
	fileSize, _ := strconv.Atoi(r.Form.Get("filesize"))
	// 2。获得redis请求链接
	rConn := pool.RedisPool().Get()
	defer func() {
		_ = rConn.Close()
	}()
	// 3。通过uploadID查询redis是否所有分块都上传完成
	data, err := redis.Values(rConn.Do("HGETALL", "FT"+uploadID))
	if err != nil {
		resp, _ := ecode.NewResponseMsg(-1, "complete upload faild", nil).JOSNBytes()
		_, _ = w.Write(resp)
	}
	totalCount := 0
	chunkCount := 0
	for i := 0; i < len(data); i++ {
		k := string(data[i].([]byte))
		v := string(data[i].([]byte))
		if k == "chunkcount" {
			totalCount, _ = strconv.Atoi(v)
		} else if strings.HasPrefix(k, "chkidx") && v == "1" {
			chunkCount += 1
		}
	}
	if totalCount != chunkCount {
		resp, _ := ecode.NewResponseMsg(-1, "invalid request", nil).JOSNBytes()
		_, _ = w.Write(resp)
		return
	}
	// 4。合并分块
	// 5。更新唯一文件表以及用户文件表
	_ = dbClient.UpdateTblFile(&models.TblFile{
		FileName: fileName,
		FileSha1: fileHash,
		FileSize: int64(fileSize),
		FileAddr: "",
		Status:   0,
		Ext1:     "",
		Ext2:     "",
	})
	_ = dbClient.UpdateTblUserFile(&models.TblUserFile{
		UserName: userName,
		FileName: fileName,
		FileSha1: fileHash,
		FileSize: int64(fileSize),
		Status:   0,
	})
	// 6。响应处理结果
	resp, _ := ecode.NewResponseMsg(0, "success", nil).JOSNBytes()
	_, _ = w.Write(resp)
}

// 取消上传分块
func CancelUploadHandler(w http.ResponseWriter, r http.Request) {

}

//查看上传分块的整体状态
func MultipartUploadStatusHandler(w http.ResponseWriter, r http.Request) {

}
