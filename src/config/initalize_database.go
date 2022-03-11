package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var defaultDatabase *xorm.Engine

func initDefaultDatabase() {
	datasource := fmt.Sprintf("%s:%s@/%s?charset=utf8", Conf.Mysql.User, Conf.Mysql.Password, Conf.Mysql.Table)
	var err error
	defaultDatabase, err = xorm.NewEngine("mysql", datasource)
	if err != nil {
		fmt.Println(err)
	}
}

func GetDefaultDatabase() *xorm.Engine {
	if defaultDatabase == nil {
		initDefaultDatabase()
	}
	return defaultDatabase
}