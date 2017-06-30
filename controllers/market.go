package controllers

import (
	"apromise-go-api/models"
	"fmt"

	"github.com/astaxie/beego"
)

type MarketController struct {
	beego.Controller
}

// @router /getList [get]
func (mc *MarketController) GetMarketList() {
	market := &models.Market{}
	markets := market.GetMarketList()
	mc.Data["json"] = markets
	fmt.Println(markets)
	mc.ServeJSON()
}
