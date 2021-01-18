package route

import (
	"github.com/feitianlove/FIleStore/handler"
	"github.com/gin-gonic/gin"
)

func NewGin() *gin.Engine {
	//gin framework，包含Logger，Recovery
	router := gin.Default()
	//处理静态资源
	router.Static("/static/", "./static")
	router.GET("/user/signup", handler.SignUpHtml)
	router.POST("/user/signup", handler.DoSignupHandler)
	//加入一个拦截器

	router.Use(handler.HTTPInterceptor2())
	return router
}
