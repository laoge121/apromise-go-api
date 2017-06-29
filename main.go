package main

import (
	"fmt"
	//"fmt"
	_ "apromise-go-api/routers"

	//"apromise-go-api/base"

	"github.com/astaxie/beego"
)

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
	fmt.Println("system start")
	beego.Run()
}
