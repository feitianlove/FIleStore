package main

import (
	"encoding/json"
	"fmt"
	"github.com/feitianlove/FIleStore/rabbitmq"
	"os"
)

//转移本地文件道ceph】

func ProcessTransfer(msg []byte) bool {
	//1。 解析msg
	pubData := rabbitmq.TransferData{}
	err := json.Unmarshal(msg, pubData)
	if err != nil {
		return false
	}
	//2。 根据临时存储文件路径，创建文件句柄
	file, err := os.Open(pubData.CurLocation)
	if err != nil {
		return false
	}
	//3。通过文件句柄将文件内容读出来并且上传到oss
	fmt.Println(file)
	//3。 更新文件的存储路径
	return true
}

//
func Run() {
	rabbitmq.Consume(rabbitmq.TransOssQueryName, "transfer_oss", ProcessTransfer)
}
