// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"apromise-go-api/controllers"

	"github.com/astaxie/beego"
	"apromise-go-api/filter"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/topic",
			beego.NSInclude(
				&controllers.TopicController{},
			),
		),
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/market",
			beego.NSInclude(
				&controllers.MarketController{},
			),
		),
		beego.NSNamespace("/topRankGoodsControl", beego.NSInclude(
			&controllers.TopRankGoodsController{},
		),
		),
	)
	beego.AddNamespace(ns)

	beego.InsertFilter("/*", beego.BeforeRouter, filter.FilterUser)


	//日志配置---我们的程序往往期望把信息输出到 log 中，现在设置输出到文件很方便，如下所示：
	//logs.SetLogger("file",`{"filename","logs/test.log"}`)
	//设置删除控制台
	//logs.BeeLogger.DelLogger("console")
	//设置日志级别
	beego.SetLevel(beego.LevelInformational)
	/*
	beego.Emergency("this is emergency")
	beego.Alert("this is alert")
	beego.Critical("this is critical")
	beego.Error("this is error")
	beego.Warning("this is warning")
	beego.Notice("this is notice")
	beego.Informational("this is informational")
	beego.Debug("this is debug")

	日志的级别如上所示的代码这样分为八个级别：

	LevelEmergency
	LevelAlert
	LevelCritical
	LevelError
	LevelWarning
	LevelNotice
	LevelInformational
	LevelDebug
	级别依次降低，默认全部打印，但是一般我们在部署环境，可以通过设置级别设置日志级别：

	beego.SetLevel(beego.LevelInformational)

	*/


	//设置输出信息
	beego.SetLogFuncCall(true)
}
