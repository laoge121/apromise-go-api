package models

import (
	//"strconv"
	//"errors"
	"fmt"

	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	TopicList map[string]*Topic
)

var (
	queryByIdSql = "SELECT id,title,market_id as marketId,sub_title as subTitle FROM t_venus_topic WHERE id =?"
)

type Topic struct {
	Id          int
	MarketId    int
	Title       string
	SubTitle    string
	Banner      string
	Description string
	AuthorId    string
	SortOrder   int
	GoodsNum    int
	ShopsNum    int
	Status      int
}

func (t Topic) String() string {

	return t.SubTitle
}

func (t *Topic) TableName() string {
	return "t_venus_topic"
}

func (t *Topic) TableEngine() string {
	return "INNODB"
}

func (t Topic) GetTopicById() (topic *Topic) {
	o := orm.NewOrm()
	o.Using("default")
	fmt.Print("<<<", t.Id)
	r := o.Raw(queryByIdSql, t.Id)
	var tmpTopic Topic
	errs := r.QueryRow(&tmpTopic)
	topic = &tmpTopic
	if errs != nil {
		fmt.Println("查询数据异常")
	}
	fmt.Println(">>", topic)
	t.Title = "雨后测试"
	r = o.Raw("update t_venus_topic set title=? where id =?", t.Title, t.Id)
	results, err := r.Exec()
	if err != nil {
		fmt.Println(">>>不ok", err)
	} else {
		fmt.Println(results.RowsAffected())
	}
	return topic
}
