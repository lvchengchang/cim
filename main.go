package main

import (
	"encoding/json"
	"log"
	"net/http"
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

func main() {
	http.HandleFunc("/user/login", userLogin)

	http.ListenAndServe(":8080", nil)
}
