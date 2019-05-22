package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"html/template"
	"log"
	"net/http"
)

var (
	DbEngin *xorm.Engine
)

type H struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func init() {
	driveName := "mysql"
	DbName := "root:lvchang@tcp(127.0.0.1:3306)/im?charset=utf8"
	DbEngin, err := xorm.NewEngine(driveName, DbName)
	if nil != err {
		log.Println(err.Error())
	}

	DbEngin.ShowSQL(true)      // show sql
	DbEngin.SetMaxOpenConns(2) // mysql max connect num

	fmt.Println("init mysql driver ok")
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
	http.Handle("/asset/", http.FileServer(http.Dir(".")))

	RegisterView()

	http.ListenAndServe(":8080", nil)
}
