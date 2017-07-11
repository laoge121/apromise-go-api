package controllers

import (
	"apromise-go-api/models"
	"fmt"

	"github.com/astaxie/beego"
	"apromise-go-api/rpc/client"
)

type MarketController struct {
	beego.Controller
}

// @router /getList [get]
func (mc *MarketController) GetMarketList() {

	client.ClientStart()

	market := &models.Market{}
	markets := market.GetMarketList()
	mc.Data["json"] = markets
	fmt.Println(markets)
	mc.ServeJSON()
}

// @router /getMarketMap [get]
func (mc *MarketController) GetMarketMap() {
	market := &models.Market{}
	res, err := market.GetMarketMap()
	if err != nil {
		mc.Data["json"] = err.Error()
	} else {
		mc.Data["json"] = res
	}
	mc.ServeJSON()
}

// @router /getLists [get]
func (mc *MarketController) GetMarketLists() {
	market := models.Market{}

	r, err := market.GetMarketLists()

	if err != nil {
		mc.Data["json"] = err.Error()
	} else {
		mc.Data["json"] = r
	}

	mc.ServeJSON()
}
