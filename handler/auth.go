package handler

import (
	"fmt"
	"github.com/feitianlove/golib/common/ecode"
	"github.com/gin-gonic/gin"
	"net/http"
)

// http 拦截器
func HTTPInterceptor(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			err := request.ParseForm()
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				fmt.Printf("parseForm err in HTTPInterceptor, err:%s\n", err)
				return
			}
			username := request.Form.Get("username")
			token := request.Form.Get("token")
			if len(username) < 3 || !IsTokenValid(token) {
				writer.WriteHeader(http.StatusForbidden)
				return
			}
			h(writer, request)
		},
	)
}

//gin 中间价
func HTTPInterceptor2() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Request.FormValue("username")
		token := c.Request.FormValue("token")
		if len(username) < 3 || !IsTokenValid(token) {
			c.Abort()
			ecode.RespErrCode(c, ecode.ErrError, "username or token is invalid")
			return
		}
		c.Next()
	}

}
