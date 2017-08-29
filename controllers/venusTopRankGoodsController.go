package controllers

import (
	"github.com/astaxie/beego"
	"apromise-go-api/models"
	"fmt"
)

type   TopRankGoodsController struct {
	beego.Controller
}

//@router /getTopRankGoodsList [get]
func (trgc *TopRankGoodsController) GetTopRankGoodsCList() {

	model := models.VenusTopRankGoodsControl{}

	model.Type = "rxdp"
	model.Status = 2
	model.Date = 1484064000

	res := model.GetTopRankGoodsControl()
	trgc.Data["json"] = res
	fmt.Println("返回结果集;{}", res)
	trgc.ServeJSON()
}