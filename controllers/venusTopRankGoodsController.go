package controllers

import (
	"github.com/astaxie/beego"
	"apromise-go-api/models"
	"fmt"
	"strconv"
	"encoding/json"
	"github.com/astaxie/beego/validation"
)

type   TopRankGoodsController struct {
	beego.Controller
}

//@router /get [get]
func (this *TopRankGoodsController) Get() {

	model := models.VenusTopRankGoodsControl{}

	model.Type = this.GetString("type")
	model.Status, _ = this.GetInt("status")
	model.Date, _ = this.GetInt("date")

	res := model.GetTopRankGoodsControl()
	this.Data["json"] = res
	fmt.Println("返回结果集;{}", res)
	this.ServeJSON()
}

//@router /post [get]
func (this *TopRankGoodsController)Post() {
	id := this.Input().Get("id")
	iid, _ := strconv.Atoi(id)
	if (iid == 0) {
		this.Ctx.WriteString(">>>>>错误")
	} else {
		this.Ctx.WriteString("<<<success>>>")
	}
}

//@router /parse [get]
func (this *TopRankGoodsController)ParseFormData() {

	//数据验证
	vali := validation.Validation{}

	model := models.VenusTopRankGoodsControl{}
	if err := this.ParseForm(&model); err != nil {
		fmt.Println(model)
		this.Ctx.WriteString("error")
	}

	beego.Informational(">>>>>>>>>>>日志测试Info>>>>>>>>>>>")
	beego.Debug(">>>>>>>>>>debug<<<<<<<<<<<<<")

	//非空校验
	res := vali.Required(model.Id, "")
	if res.Ok != true {

		fmt.Println(">>>>>???", res)

		//数据异常那么就终止继续执行,遇到下面的状态码 那么停止继续往下执行
		this.Data["content"] = ">>>>>>>>>>>>>>>>"
		this.Abort("404")
	}

	//
	fmt.Println(model)
	this.Ctx.WriteString("success")
}




//@router /json [post]
func (this *TopRankGoodsController)PostJsons() {

	var ob models.VenusTopRankGoodsControl

	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)

	fmt.Println(ob)
	b, _ := json.Marshal(ob);
	this.Data["json"] = string(b)
	this.ServeJSON()

}







