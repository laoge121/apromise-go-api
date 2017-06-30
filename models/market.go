package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Market struct {
	Id     int
	Title  string
	Date   string
	Status int
}

func (m Market) String() string {
	return m.Title
}

func (m *Market) TableName() string {
	return "t_venus_market"
}

func (m *Market) TableEngine() string {
	return "INNODB"
}

func (m *Market) GetMarketList() *[]Market {
	o := orm.NewOrm()

	r := o.Raw("SELECT id,title,date,status from t_venus_market ")

	var mk []Market

	num, err := r.QueryRows(&mk)
	if nil != err {
		fmt.Println("market info line num:", num)
	}

	return &mk
}
