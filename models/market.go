package models

import (
	"errors"
	"fmt"

	"github.com/astaxie/beego/orm"
)

var (
	market_getMarketList = "SELECT id,title,date,status from t_venus_market"

	market_getMarket = "SELECT * FROM t_venus_market where status=?"
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

	r := o.Raw(market_getMarketList)

	var mk []Market

	num, err := r.QueryRows(&mk)
	if nil != err {
		fmt.Println("market info line num:", num)
	}

	return &mk
}

func (m *Market) GetMarketMap() (*[]orm.Params, error) {
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw(market_getMarket, 1).Values(&maps)
	if err != nil {
		err = errors.New("查询数据异常！")
	}
	fmt.Print("返回数据行数：", num, "\n")
	return &maps, err
}

func (m *Market) GetMarketLists() (*[]orm.ParamsList, error) {
	o := orm.NewOrm()
	var lists []orm.ParamsList
	num, err := o.Raw(market_getMarketList).ValuesList(&lists)
	if err != nil && num > 0 {
		err = errors.New("查询数据异常")
	}
	fmt.Println("返回数据集:", num)
	return &lists, err
}
