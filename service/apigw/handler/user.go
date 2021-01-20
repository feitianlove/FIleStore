package handler

import (
	"context"
	"github.com/feitianlove/FIleStore/service/account/proto"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"net/http"
)

var (
	userClient proto.UserService
)

func init() {
	//	service := micro.NewService(micro.Name("go.micro.api.user")) 如果没有micro.Name()表示不注册到注册中心
	service := micro.NewService()
	//  解析命令行参数
	service.Init()
	//  初始化rpcClient
	userClient = proto.NewUserService("go.micro.service.user", service.Client())
}

// 使用grpc
func DoSignupHandler(c *gin.Context) {
	username := c.Request.FormValue("username")
	passwd := c.Request.FormValue("passwd")

	resp, err := userClient.Signup(context.TODO(), &proto.ReqSignup{
		Username: username,
		Passwd:   passwd,
	})
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Message,
	})
}

func SignupHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/view/signup.html")
}
