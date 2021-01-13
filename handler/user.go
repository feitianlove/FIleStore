package handler

import (
	"fmt"
	"github.com/feitianlove/FIleStore/models"
	"github.com/feitianlove/FIleStore/util"
	"github.com/feitianlove/golib/common/ecode"
	"github.com/feitianlove/golib/common/predicate"
	"io/ioutil"
	"net/http"
	"time"
)

// 盐
const pwdSalt = "ftfeng"

//注册
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
	if len(username) < 3 || len(passwd) < 4 {
		_, _ = w.Write([]byte("invalid parameter"))
		return
	}
	encPasswd := util.Sha1([]byte(passwd + pwdSalt))
	user := models.TblUser{
		UserName:   username,
		UserPwd:    encPasswd,
		LastActive: time.Now(),
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

// 登陆
func SignInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data, err := ioutil.ReadFile("./static/view/signin.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("opne ./static/view/signin.html err, err:%s\n", err)
			return
		}
		_, _ = w.Write(data)
		return
	}
	//1。校验用户名和密码
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("parseForm err in SignupHandler, err:%s\n", err)
		return
	}
	username := r.Form.Get("username")
	passwd := r.Form.Get("passwd")
	useMessage, err := dbClient.GetTblUserByUserName(predicate.EqualPredicate, username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("error"))
		return
	}
	if useMessage[0].UserPwd != util.Sha1([]byte(passwd+pwdSalt)) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("用户名密码错误"))
		return
	}
	//2。生成访问凭证
	token := GetToken(username)
	err = dbClient.CreateTblUserToken(&models.TblUserToken{
		UserName:  username,
		UserToken: token,
	})
	if useMessage[0].UserPwd != util.Sha1([]byte(passwd+pwdSalt)) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("初始化token error"))
		return
	}
	//3。登陆成功后重定向到首页,返回 url 路径前端跳转
	//http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	msg := ecode.NewResponseMsg(200, "", struct {
		Token string
	}{Token: token})
	data, _ := msg.JOSNBytes()
	_, _ = w.Write(data)

}

// 生成token
func GetToken(userName string) string {
	//md5(username+timestamp+token_salt) + timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokePrefix := util.MD5([]byte(userName + ts + "token_salt"))
	return tokePrefix + ts[:8]
}

// 获取用户信息
func UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	// 1. 解析用户请求参数
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("parseForm err in SignupHandler, err:%s\n", err)
		return
	}
	username := r.Form.Get("username")
	token := r.Form.Get("token")
	// 2. 验证token是否有效
	if !IsTokenValid(token) {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte("error"))
		return
	}
	// 3. 查询用户信息
	useInfo, err := dbClient.GetTblUserByUserName(predicate.EqualPredicate, username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("error"))
		return
	}
	// 4. 组装并且响应用户数据
	msg := ecode.NewResponseMsg(200, "", struct {
		data interface{}
	}{data: useInfo[0]})
	data, _ := msg.JOSNBytes()
	_, _ = w.Write(data)
}

// 判断token是否有效
func IsTokenValid(token string) bool {
	//TODO
	// 1. 判断token 的时效性，是否过期
	// 2. 查询是否有username和token的信息，是否一致
	return true
}
