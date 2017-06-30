package models

import (
	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
)

func init() {
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpass")
	url := beego.AppConfig.String("mysqlurls")
	db := beego.AppConfig.String("mysqldb")

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", db, user+":"+pwd+"@"+url+"?charset=utf8", 30)
}
