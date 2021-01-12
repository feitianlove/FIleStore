package models

import "time"

type TblFile struct {
	FileName string
	FileSha1 string
	FileSize int64
	FileAddr string
}

// 用户表
type TblUser struct {
	Id             int
	UserName       string
	UserPwd        string
	Email          string
	Phone          string
	EmailValidated int8
	PhoneValidate  int8
	SignUP         time.Time
	LastActive     time.Time
	Profile        string
	Status         int
}

// 用户鉴权表

type TblUserToken struct {
	Id        int
	UserName  string
	UserToken string
}
