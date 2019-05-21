package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/user/login",
		func(w http.ResponseWriter, r *http.Request) {
			// mysql 逻辑处理
			r.ParseForm()

			mobile := r.PostForm.Get("mobile")
			passwd := r.PostForm.Get("passwd")

			loginok := false
			if mobile == "11647791" && passwd == "123" {
				loginok = true
			}

			if loginok {
				str := `{"code":0,"msg":"ok","data":{"id":1,"token":"ok"}}`
				w.Header().Set("Content-Type", "application-json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(str))
				return
			}
			str := `{"code":6009,"msg":"fail","data":"fail"}`
			w.Header().Set("Content-Type", "application-json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(str))
		})

	http.ListenAndServe(":8080", nil)
}
