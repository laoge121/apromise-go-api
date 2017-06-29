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

	return t.Title
}

func (t *Topic) TableName() string {
	return "t_venus_topic"
}

func (t *Topic) TableEngine() string {
	return "INNODB"
}

func init() {
	TopicList = make(map[string]*Topic)
	t := Topic{1, 2, "标题", "副标题", "标题图", "描述", "雨后", 1, 12, 23, 1}
	TopicList[t.Title] = &t
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "test:test123@tcp(10.13.1.15:3306)/venus?charset=utf8", 30)
}

func (t Topic) GetTopicById() {
	o := orm.NewOrm()
	o.Using("default")

	r := o.Raw("SELECT title FROM t_venus_topic WHERE id =?", t.Id)
	var topic Topic
	errs := r.QueryRow(&topic)
	if errs != nil {
		fmt.Println("查询数据异常")
	}
	fmt.Println(topic)

	r = o.Raw("update t_venus_topic set title=? where id =?", t.Title, t.Id)
	results, err := r.Exec()
	if err != nil {
		fmt.Println(">>>不ok", err)
	} else {
		fmt.Println(results.RowsAffected())
	}
}
