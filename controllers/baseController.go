package controllers

import (
	"github.com/astaxie/beego"
	"time"
)

type NestPreparer interface {
	NestPrepare()
}

type baseRouter struct {
	beego.Controller
	isLogin bool
}

func (this *baseRouter) Prepare() {
	this.Data["PageStartTime"] = time.Now()
	/*this.Data["AppKeywords"] = utils.AppKeywords
	this.Data["AppName"] = utils.AppName
	this.Data["AppVer"] = utils.AppVer
	this.Data["AppUrl"] = utils.AppUrl
	this.Data["AppLogo"] = utils.AppLogo
	this.Data["AvatarURL"] = utils.AvatarURL
	this.Data["IsProMode"] = utils.IsProMode*/

	if app, ok := this.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}

