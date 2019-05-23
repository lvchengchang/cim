package controller

import (
	"cim/model"
	"cim/service"
	"cim/util"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

var (
	userService service.UserService
)

func UserLogin(w http.ResponseWriter, r *http.Request) {
	// mysql 逻辑处理
	r.ParseForm()

	mobile := r.PostForm.Get("mobile")
	passwd := r.PostForm.Get("passwd")

	tmp := model.User{}
	_, err := service.DbEngin.Where("mobile= ? ", mobile).Get(&tmp)
	if err != nil {
		util.RespFail(w, err.Error())
		return
	}

	if tmp.Passwd == util.MD5Encode(passwd) {
		util.Resp(w, 0, "ok", tmp)
		tmp.Token = util.MD5Encode(fmt.Sprintf("%8d", rand.Intn(8)))
		_, err := service.DbEngin.ID(tmp.Id).Cols("token").Update(&tmp)
		if err != nil {
			log.Fatalln(err)
		}

		return
	}

	util.RespFail(w, "passwd error")
	return
}

func UserRegister(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	mobile := r.PostForm.Get("mobile")
	passwd := r.PostForm.Get("passwd")
	nickname := fmt.Sprintf("user%6d", rand.Intn(31))
	avatar := ""
	sex := model.SEX_MEN
	user, err := userService.Register(mobile, passwd, nickname, avatar, sex)
	if err != nil {
		util.Resp(w, -1, err.Error(), "fail")
		return
	}

	util.Resp(w, 0, "", user)
}
