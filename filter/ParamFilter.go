package filter

import (
	"fmt"
	"github.com/astaxie/beego/context"
)

var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("uid").(int)

	fmt.Println(">>>>>>FilterUser>>>>>>>")

	if ok {
		fmt.Println("<<>><>成功!<><")
	}
	/*
		if !ok && ctx.Request.RequestURI != "/login" {
			ctx.Redirect(302, "login")
		}*/
}
