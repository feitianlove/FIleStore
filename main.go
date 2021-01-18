package main

import "github.com/feitianlove/FIleStore/route"

func main() {
	// TODO 普通的net http
	//http.HandleFunc("/file/upload", handler.UploadHandler)
	//http.HandleFunc("/file/upload/suc", handler.UpLoadSucHandler)
	//http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	//http.HandleFunc("/file/download", handler.DownloadHandler)
	//http.HandleFunc("/file/update", handler.FileMetaUpdateHandler)
	//http.HandleFunc("/file/delete", handler.FileDeleteHandler)
	//http.HandleFunc("/user/signup", handler.SignupHandler)
	//http.HandleFunc("/user/signin", handler.SignInHandler)
	////http.HandleFunc("/user/info", handler.UserInfoHandler)
	//// 加了拦截器
	//http.HandleFunc("/user/info", handler.HTTPInterceptor(handler.UserInfoHandler))
	//
	//err := http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	panic(err)
	//}

	// GIN
	router := route.NewGin()
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}

}

//
