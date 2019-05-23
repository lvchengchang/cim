package service

import (
	"cim/model"
	"cim/util"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type UserService struct {
}

func (s *UserService) Register(mobile,
	passwd,
	nickname,
	avatar,
	sex string) (user model.User, err error) {
	// user is exist
	tmp := model.User{}
	_, err = DbEngin.Where("mobile=? ", mobile).Get(&tmp)
	if err != nil {
		return tmp, err
	}

	if tmp.Id > 0 {
		return tmp, errors.New("user is exist")
	}

	// insert value
	tmp.Mobile = mobile
	tmp.Avatar = avatar
	tmp.Nickname = nickname
	tmp.Passwd = util.MD5Encode(passwd)
	tmp.Createat = time.Now()
	tmp.Sex = sex
	tmp.Token = util.MD5Encode(fmt.Sprintf("%8d", rand.Intn(8)))

	_, err = DbEngin.InsertOne(&tmp)
	if err != nil {
		return tmp, nil
	}

	return tmp, err
}

func (s *UserService) Login(mobile,
	passwd string) (user model.User, err error) {
	return user, nil
}
