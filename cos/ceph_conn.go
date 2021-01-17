package cos

import (
	"gopkg.in/amz.v1/aws"
	"gopkg.in/amz.v1/s3"
)

var S3Conn *s3.S3

func NewS3Conection() *s3.S3 {
	if S3Conn != nil {
		return S3Conn
	}
	// 1。 初始化对象存储接口信息
	auth := aws.Auth{
		AccessKey: "",
		SecretKey: "",
	}
	curRegion := aws.Region{
		Name:                 "default",
		EC2Endpoint:          "http://127.0.0.1:9080",
		S3Endpoint:           "http://127.0.0.1:9080",
		S3BucketEndpoint:     "",
		S3LocationConstraint: false,
		S3LowercaseBucket:    false,
		Sign:                 aws.SignV2,
	}
	// 2。创建s3类型的链接
	return s3.New(auth, curRegion)
}

// 获取指定bucket对象
func GetS3Bucket(bucket string) *s3.Bucket {
	conn := NewS3Conection()
	return conn.Bucket(bucket)
}
