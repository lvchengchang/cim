package util

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

func RespFail(w http.ResponseWriter, msg string) {
	Resp(w, -1, msg, nil)
}

func RespOk(w http.ResponseWriter, msg string, data interface{}) {
	Resp(w, 0, msg, data)
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
