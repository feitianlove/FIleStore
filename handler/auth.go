package handler

import (
	"fmt"
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
