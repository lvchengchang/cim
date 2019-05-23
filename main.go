package main

import (
	"cim/model"
	"cim/service"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"math/rand"
	"net/http"
)

var (
	userService service.UserService
)

type H struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func userLogin(w http.ResponseWriter, r *http.Request) {
	// mysql 逻辑处理
	r.ParseForm()

	mobile := r.PostForm.Get("mobile")
	passwd := r.PostForm.Get("passwd")

	loginok := false
	if mobile == "11647791" && passwd == "123" {
		loginok = true
	}

	if loginok {
		data := make(map[string]interface{})
		data["id"] = 1
		data["token"] = "test"
		Resp(w, 0, "ok", data)
		return
	}

	Resp(w, 0, "fail", "fail")
	return
}

func register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	mobile := r.PostForm.Get("mobile")
	passwd := r.PostForm.Get("passwd")
	nickname := fmt.Sprintf("user%6d", rand.Intn(31))
	avatar := ""
	sex := model.SEX_MEN
	user, err := userService.Register(mobile, passwd, nickname, avatar, sex)
	if err != nil {
		Resp(w, -1, err.Error(), "fail")
		return
	}

	Resp(w, 0, "", user)
}

func Resp(w http.ResponseWriter, code int, msg string, data interface{}) {
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code: code,
		Msg:  msg,
		Data: data,
	}

	str, err := json.Marshal(h)
	if err != nil {
		log.Fatal(err)
		return
	}
	// define struct
	w.Write([]byte(str))
}

func RegisterView() {
	// analyze ** 表示目录 * 表示文件
	tpl, err := template.ParseGlob("view/**/*")
	if err != nil {
		log.Fatalln(err.Error())
	}
	for _, v := range tpl.Templates() {
		tplname := v.Name()

		http.HandleFunc(tplname, func(w http.ResponseWriter, r *http.Request) {
			tpl.ExecuteTemplate(w, tplname, nil)
		})
	}
}

func main() {
	http.HandleFunc("/user/login", userLogin)
	http.HandleFunc("/user/register", register)

	http.Handle("/asset/", http.FileServer(http.Dir(".")))

	RegisterView()

	http.ListenAndServe(":8080", nil)
}
