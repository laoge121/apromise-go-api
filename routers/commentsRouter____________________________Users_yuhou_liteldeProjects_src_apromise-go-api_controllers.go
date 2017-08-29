package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["apromise-go-api/controllers:MarketController"] = append(beego.GlobalControllerRouter["apromise-go-api/controllers:MarketController"],
		beego.ControllerComments{
			Method: "GetMarketList",
			Router: `/getList`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apromise-go-api/controllers:MarketController"] = append(beego.GlobalControllerRouter["apromise-go-api/controllers:MarketController"],
		beego.ControllerComments{
			Method: "GetMarketMap",
			Router: `/getMarketMap`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apromise-go-api/controllers:MarketController"] = append(beego.GlobalControllerRouter["apromise-go-api/controllers:MarketController"],
		beego.ControllerComments{
			Method: "GetMarketLists",
			Router: `/getLists`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apromise-go-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["apromise-go-api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apromise-go-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["apromise-go-api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apromise-go-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["apromise-go-api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apromise-go-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["apromise-go-api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apromise-go-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["apromise-go-api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apromise-go-api/controllers:TopRankGoodsController"] = append(beego.GlobalControllerRouter["apromise-go-api/controllers:TopRankGoodsController"],
		beego.ControllerComments{
			Method: "GetTopRankGoodsCList",
			Router: `/getTopRankGoodsList`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apromise-go-api/controllers:TopicController"] = append(beego.GlobalControllerRouter["apromise-go-api/controllers:TopicController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apromise-go-api/controllers:UserController"] = append(beego.GlobalControllerRouter["apromise-go-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apromise-go-api/controllers:UserController"] = append(beego.GlobalControllerRouter["apromise-go-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apromise-go-api/controllers:UserController"] = append(beego.GlobalControllerRouter["apromise-go-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apromise-go-api/controllers:UserController"] = append(beego.GlobalControllerRouter["apromise-go-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apromise-go-api/controllers:UserController"] = append(beego.GlobalControllerRouter["apromise-go-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apromise-go-api/controllers:UserController"] = append(beego.GlobalControllerRouter["apromise-go-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apromise-go-api/controllers:UserController"] = append(beego.GlobalControllerRouter["apromise-go-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
