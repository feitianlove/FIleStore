package handler

import (
	"context"
	"github.com/feitianlove/FIleStore/config"
	"github.com/feitianlove/FIleStore/models"
	"github.com/feitianlove/FIleStore/service/account/proto"
	"github.com/feitianlove/FIleStore/store"
	"github.com/feitianlove/FIleStore/util"
	"time"
)

type User struct {
}

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

func (u *User) Signup(ctx context.Context, req *proto.ReqSignup, res *proto.RespSignup) error {
	username := req.Username
	passwd := req.Passwd
	if len(username) < 3 || len(passwd) < 4 {
		res.Code = -1
		res.Message = "注册参数无效"
		return nil
	}
	encPasswd := util.Sha1([]byte(passwd + "ftfeng"))
	user := models.TblUser{
		UserName:   username,
		UserPwd:    encPasswd,
		Phone:      "679",
		LastActive: time.Now(),
	}
	err := dbClient.CreateTblUser(&user)
	if err != nil {
		res.Code = -1
		res.Message = "注册参数失败"
		return nil
	}
	res.Code = 0
	res.Message = "ok"
	return nil
}
