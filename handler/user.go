package handler

import (
	"fmt"
	"github.com/feitianlove/FIleStore/models"
	"github.com/feitianlove/FIleStore/util"
	"io/ioutil"
	"net/http"
	"time"
)

// 盐
const pwdSalt = "ftfeng"

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("opne ./static/view/signup.html err, err:%s\n", err)
			return
		}
		_, _ = w.Write(data)
		return
	}
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("parseForm err in SignupHandler, err:%s\n", err)
		return
	}
	username := r.Form.Get("username")
	passwd := r.Form.Get("passwd")
	fmt.Println(username, passwd)
	if len(username) < 3 || len(passwd) < 4 {
		_, _ = w.Write([]byte("invalid parameter"))
		return
	}
	encPasswd := util.Sha1([]byte(passwd + pwdSalt))
	user := models.TblUser{
		UserName: username,
		UserPwd:  encPasswd,
	}
	err = dbClient.CreateTblUser(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("error"))
		fmt.Println(err)
		return
	}
	_, _ = w.Write([]byte("success"))
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	//1。校验用户名和密码

	username := ""
	//2。生成访问凭证
	token := GetToken(username)
	fmt.Println(token)
	//3。登陆成功后重定向到首页,返回 url 路径前端跳转

}

func GetToken(userName string) string {
	//md5(username+timestamp+token_salt) + timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokePrefix := util.MD5([]byte(userName + ts + "token_salt"))
	return tokePrefix + ts[:8]
}
