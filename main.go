package main

import (
	//"fmt"
	//"fmt"
	_ "apromise-go-api/routers"

	//"apromise-go-api/base"

	//"apromise-go-api/models"

	"github.com/astaxie/beego"

	"apromise-go-api/filter"
)

func init(){

	beego.InsertFilter("/*", beego.BeforeRouter, filter.FilterUser)


	//日志配置---我们的程序往往期望把信息输出到 log 中，现在设置输出到文件很方便，如下所示：
	//logs.SetLogger("file",`{"filename","logs/test.log"}`)
	//设置删除控制台
	//logs.BeeLogger.DelLogger("console")
	//设置日志级别
	//beego.SetLevel(beego.LevelInformational)
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

	*/
	beego.SetLevel(beego.LevelDebug)



	//设置输出信息
	beego.SetLogFuncCall(true)
}

func baseTest() {
	/*if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
	*/
	/*base.VarTest()
	base.VarTypeChange()
	base.ConstCh()
	base.PointT()
	base.IfElseT()

	base.ArratT()
	base.PxMp()

	base.SliceT()
	base.MapT()*/

	//var arrys = [...]int{1, 6, 2, 3, 4}
	//slis := arrys[:len(arrys)]
	//base.FuncT(slis)

	//匿名函数
	/*a := 10
	b := base.Change(a)
	fmt.Println(b(3))
	fmt.Println(b(2))
	fmt.Println(a)

	base.DeferT()

	base.StructT()
	*/

	/*A := base.A{}
	B := base.B{}
	A.Println()
	A.Println1()
	B.Println()
	B.Println2()

	(*base.A).Println(&A)
	var t1 base.IntTr
	t1.Increment()
	fmt.Println(t1)
	*/
	/*var ph base.Usb
	phone := base.PhoneConnector{"dianhua"}
	ph = phone
	base.DisConnect(ph)

	tv := base.TvConnector("tv")
	tv.Connect()
	var con base.Connecter
	con = tv
	base.DisConnectInf(con)


	base.ThreadT()
	fmt.Println(">>>>>>>>>>>>.")
	base.GoBockT()
	base.GoThread()
	base.GoSys()
	base.SelectT()
	base.SelectT1()
	*/
}

func main() {

	//fmt.Println("system start")
	//go service.Start()
	beego.Run()
	//topic := models.Topic{}
	//topic.Id = 3580
	//topic.GetTopicById()
}
