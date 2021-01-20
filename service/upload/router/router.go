package router

import (
	"github.com/feitianlove/FIleStore/service/upload/api"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.Static("/static/", "./static")

	//TODO 接口未实行
	router.GET("/user/signup", api.DoUploadHandlerGen)
	router.POST("/user/signup", api.DoUploadHandlerGen)
	return router
}
