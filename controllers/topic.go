package controllers

import (
	"apromise-go-api/models"
	"strconv"

	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

// @router /:id [get]
func (tc *TopicController) Get() {

	id := tc.GetString(":id")

	if id != "" {
		topic := models.Topic{}
		topic.Id, _ = strconv.Atoi(id)
		tmpTopic := topic.GetTopicById()
		tc.Data["json"] = tmpTopic
	} else {
		tc.Data["json"] = "查询错误"
	}
	tc.ServeJSON()
}
