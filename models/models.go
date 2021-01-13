package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TblFile struct {
	gorm.Model
	FileName string `json:"file_name"`
	FileSha1 string `json:"file_sha_1"`
	FileSize int64  `json:"file_size"`
	FileAddr string `json:"file_addr"`
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
	SignUp         time.Time `json:"sign_up"`
	LastActive     time.Time `json:"last_active"`
	Profile        string    `json:"profile" gorm:"type:text"`
	Status         int       `json:"status"`
}

// 用户鉴权表
type TblUserToken struct {
	gorm.Model
	UserName  string `json:"user_name" gorm:"type:varchar(64)"`
	UserToken string `json:"user_token" gorm:"type:varchar(256)"`
}
