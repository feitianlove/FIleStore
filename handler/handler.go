package handler

import (
	"encoding/json"
	"fmt"
	"github.com/feitianlove/FIleStore/config"
	"github.com/feitianlove/FIleStore/meta"
	"github.com/feitianlove/FIleStore/models"
	"github.com/feitianlove/FIleStore/store"
	"github.com/feitianlove/FIleStore/util"

	"io"
	"io/ioutil"
	"net/http"
	"os"
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
