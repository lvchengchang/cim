package service

import (
	"cim/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var (
	DbEngin *xorm.Engine
)

func init() {
	driveName := "mysql"
	DbName := "root:lvchang@tcp(118.25.180.168:3306)/im?charset=utf8"
	var err error
	DbEngin, err = xorm.NewEngine(driveName, DbName)
	if nil != err {
		log.Println(err.Error())
	}

	DbEngin.ShowSQL(true)      // show sql
	DbEngin.SetMaxOpenConns(2) // mysql max connect num

	DbEngin.Sync2(new(model.User))

	fmt.Println("init mysql driver ok")
}
