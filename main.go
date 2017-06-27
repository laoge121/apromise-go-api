package main

import (
	"fmt"
	//_ "apromise-go-api/routers"

	"apromise-go-api/base"

	//"github.com/astaxie/beego"
)

func main() {
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
	a := 10
	b := base.Change(a)
	fmt.Println(b(3))
	fmt.Println(b(2))
	fmt.Println(a)

	base.DeferT()
}
