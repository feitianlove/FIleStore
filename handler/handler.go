package handler

import (
	"encoding/json"
	"fmt"
	"github.com/feitianlove/FIleStore/config"
	"github.com/feitianlove/FIleStore/cos"
	"github.com/feitianlove/FIleStore/meta"
	"github.com/feitianlove/FIleStore/models"
	"github.com/feitianlove/FIleStore/store"
	"github.com/feitianlove/FIleStore/util"
	"github.com/feitianlove/golib/common/ecode"
	"github.com/feitianlove/golib/common/predicate"
	"gopkg.in/amz.v1/s3"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

//初始化db
var (
	dbClient *store.Store
)

func init() {
	conf := config.Config{MysqlConf: &config.MysqlConf{
		User:     "root",
		Pass:     "ftfeng123",
		Host:     "127.0.0.1",
		Port:     3306,
		Database: "fileserver",
	}}
	c, err := store.NewStore(&conf)
	if err != nil {
		panic(err)
	}
	dbClient = c

}

// 处理文件上传
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			_, _ = io.WriteString(w, "inner error")
			return
		}
		_, _ = io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		file, header, err := r.FormFile("file")
		defer func() {
			_ = file.Close()
		}()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("Fail  to get data,err %s\n", err)
			return
		}

		newFile, err := os.Create("/tmp/" + header.Filename)
		defer func() {
			_ = newFile.Close()
		}()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("Fail create file,err %s\n", err)
			return
		}
		// 添加文件元信息
		fileMeta := meta.FileMeta{
			FileSha1: "",
			FileName: header.Filename,
			FileSize: 0,
			Location: "/tmp/" + header.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:05:05"),
		}
		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("Fail to save data into file,err %s\n", err)
			return
		}
		// TODO 将文件保存到ceph中
		_, _ = newFile.Seek(0, 0)
		data, _ := ioutil.ReadAll(newFile)
		bucket := cos.GetS3Bucket("userFile")
		cephPath := "/ceph/" + fileMeta.FileSha1
		err = bucket.Put(cephPath, data, "octet-stream", s3.PublicRead)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("Fail to save data into ceph,err %s\n", err)
			return
		}
		fileMeta.Location = cephPath
		//TODO 计算file hash
		_, _ = newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		// 文件中保存
		meta.UpdateFileMeta(fileMeta)
		// v2 存储db
		err = dbClient.CreateTblFile(&models.TblFile{
			FileName: fileMeta.FileName,
			FileSha1: fileMeta.FileSha1,
			FileSize: fileMeta.FileSize,
			FileAddr: fileMeta.Location,
		})
		if err != nil {
			fmt.Println(err)
		}
		// TODO 更新文件表
		err = r.ParseForm()
		if err != nil {
			fmt.Printf("Fail to parse form, err %s ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		username := r.Form.Get("username")
		err = dbClient.CreateTblUserFile(&models.TblUserFile{
			UserName: username,
			FileName: fileMeta.FileName,
			FileSha1: fileMeta.FileSha1,
			FileSize: fileMeta.FileSize,
			Status:   0,
		})
		if err != nil {
			_, _ = w.Write([]byte("upload Failed"))
		}
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

// 提示用户上传成功
func UpLoadSucHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "upload finish")
}

// 查询文件信息
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("Fail to parse form, err %s ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fileHash := r.Form["fileHash"][0]
	fMeta := meta.GetFileMeta(fileHash)
	data, err := json.Marshal(fMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(data)
}

//根据查询用户的文件信息
func FileQueryHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("Fail to parse form, err %s ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	limitCut, _ := strconv.Atoi(r.Form.Get("limit"))
	username := r.Form.Get("username")
	data, err := dbClient.GetTblUserFileByUserName(predicate.EqualPredicate, username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("query file by username err"))
	}
	mes, _ := json.Marshal(data[:limitCut])
	_, _ = w.Write(mes)
}

// 文件下载
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("Fail to parse form, err %s ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fsha1 := r.Form.Get("fileHash")
	fm := meta.GetFileMeta(fsha1)
	f, err := os.Open(fm.Location)
	defer func() {
		_ = f.Close()
	}()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//TODO 这里都是小文件，如果文件比较大，就需要实现一个流的模式
	data, err := ioutil.ReadAll(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/octect-stream")
	w.Header().Set("Content-disposition", "attachment;filename=\""+fm.FileName+"\"")
	_, _ = w.Write(data)
}

// 更新文件元信息接口
func FileMetaUpdateHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("Fail to parse form, err %s ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	opType := r.Form.Get("op")
	if opType == "0" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fileSha1 := r.Form.Get("fileHash")
	newFileName := r.Form.Get("fileName")
	curFileMeta := meta.GetFileMeta(fileSha1)
	curFileMeta.FileName = newFileName
	meta.UpdateFileMeta(curFileMeta)
	data, err := json.Marshal(curFileMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

// 文件删除的接口
func FileDeleteHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("Fail to parse form, err %s ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fileSha1 := r.Form.Get("fileHash")
	fMeta := meta.GetFileMeta(fileSha1)
	//先删除索引
	meta.RemoveFileMeta(fileSha1)
	// 在删除文件
	err = os.Remove(fMeta.Location)
	if err != nil {
		fmt.Printf("Fail delete file %s ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// 妙传逻辑
func TryFastUpLoadHandler(w http.ResponseWriter, r *http.Request) {
	// 1。解析请求参数
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("Fail to parse form, err %s ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fileSha1 := r.Form.Get("fileHash")
	userName := r.Form.Get("username")
	fileName := r.Form.Get("filename")
	fileSize, _ := strconv.Atoi(r.Form.Get("filesize"))
	// 2。从文件表中查询相同hash的文件记录
	fileMeta, err := dbClient.GetTblFileByFileSha1(predicate.EqualPredicate, fileSha1)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 3。查不到记录则返回妙传失败
	if len(fileMeta) == 0 {
		resp := &ecode.RespMsg{
			Code: -1,
			Msg:  "妙传失败，请访问普通上传接口",
			Data: nil,
		}
		//resp = ecode.NewResponseMsg(-1, "妙传失败，请访问普通上传接口", nil)
		result, _ := resp.JOSNBytes()
		_, _ = w.Write(result)
	}
	// 4。上传过则将文件信息写入用户文件表，返回成功
	err = dbClient.CreateTblUserFile(&models.TblUserFile{
		UserName: userName,
		FileName: fileName,
		FileSha1: fileSha1,
		FileSize: int64(fileSize),
		Status:   0,
	})
	if err != nil {
		_, _ = w.Write([]byte("upload Failed"))
	}
	resp := &ecode.RespMsg{
		Code: 0,
		Msg:  "success",
		Data: nil,
	}
	//resp = ecode.NewResponseMsg(-1, "妙传失败，请访问普通上传接口", nil)
	result, _ := resp.JOSNBytes()
	_, _ = w.Write(result)
}

//TODO 相同文件冲突处理
// 1。 允许不同用户同时上传同一个文件
// 2。先完成的先入库
// 3。后面上传的只更新用户文件表，并删除已经上传的文件表
