package main

import (
	"cim/controller"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)

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
	http.HandleFunc("/user/login", controller.UserLogin)
	http.HandleFunc("/user/register", controller.UserRegister)

	http.Handle("/asset/", http.FileServer(http.Dir(".")))

	RegisterView()

	http.ListenAndServe(":8080", nil)
}
