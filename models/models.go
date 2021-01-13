package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TblFile struct {
	gorm.Model
	FileName string `json:"file_name"  gorm:"type:varchar(256)"`
	FileSha1 string `json:"file_sha_1"  gorm:"type:varchar(64) unique"`
	FileSize int64  `json:"file_size"  gorm:"type:varchar(20)"`
	FileAddr string `json:"file_addr"  gorm:"type:varchar(1024)"`
	Status   int    `json:"status" gorm:"index"`
	Ext1     string `json:"ext_1" gorm:"type:int"`
	Ext2     string `json:"ext_2" gorm:"type:text"`
}

// 用户表
type TblUser struct {
	gorm.Model
	UserName       string    `json:"user_name" gorm:"type:varchar(64)"`
	UserPwd        string    `json:"user_pwd" gorm:"type:varchar(256)"`
	Email          string    `json:"email" gorm:"type:varchar(128)"`
	Phone          string    `json:"phone" gorm:"unique" sql:"type: varchar(128)"`
	EmailValidated int       `json:"email_validated" gorm:"type:tinyint(1)"`
	PhoneValidate  int       `json:"phone_validate" gorm:"type:tinyint(1)"`
	LastActive     time.Time `json:"last_active"`
	Profile        string    `json:"profile" gorm:"type:text"`
	Status         int       `json:"status" gorm:"index"`
}

// 用户鉴权表
type TblUserToken struct {
	gorm.Model
	UserName  string `json:"user_name" gorm:"unique type:varchar(64)" `
	UserToken string `json:"user_token" gorm:"type:varchar(256)"`
}

type TblUserFile struct {
	gorm.Model
	UserName string `json:"user_name" gorm:"index;type:varchar(64);unique_index:S_R"`
	FileName string `json:"file_name"  gorm:"type:varchar(256)"`
	FileSha1 string `json:"file_sha_1"  gorm:"type:varchar(64);unique_index:S_R"`
	FileSize int64  `json:"file_size"  gorm:"type:varchar(20)"`
	Status   int    `json:"status" gorm:"index"`
}
