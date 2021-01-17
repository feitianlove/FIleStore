package cos

import (
	"fmt"
	"gopkg.in/amz.v1/s3"
	"testing"
)

func TestNewS3Conection(t *testing.T) {
	bucket := GetS3Bucket("test")
	// 创建一个新的bucket
	err := bucket.PutBucket(s3.PublicRead)
	if err != nil {
		panic(err)
	}
	// 查询这个bucket下面指定条件的object keys
	res, err := bucket.List("", "", "", 100)
	fmt.Printf("object keys:%+v\n", res)
	// 新上传一个对象
	err = bucket.Put("./test", []byte("just for test"), "octet-stream", s3.PublicRead)
	if err != nil {
		panic(err)
	}
}
