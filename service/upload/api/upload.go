package api

import (
	"github.com/feitianlove/FIleStore/cos"
	"github.com/feitianlove/FIleStore/meta"
	"github.com/feitianlove/FIleStore/models"
	"github.com/feitianlove/FIleStore/store"
	"github.com/feitianlove/FIleStore/util"
	"gopkg.in/amz.v1/s3"
	"io"

	"io/ioutil"
	"time"

	"github.com/feitianlove/FIleStore/config"

	"fmt"
	"net/http"
	"os"
)

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
func DoUploadHandler(w http.ResponseWriter, r *http.Request) {
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

}
