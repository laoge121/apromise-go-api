package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"errors"
)

var (
	queryTopRankList = "SELECT id,date,type,stime,etime,status,create_id,update_id,created,updated from t_venus_top_rank_goods_control where date=? and type =?  and status=2 limit 10 "

	updateTopRank = "UPDATE t_venus_top_rank_goods_control SET update_id=? where id=? limit 1"
)

type VenusTopRankGoodsControl struct {
	Id        int   `form:"id"`      // '自增主键',
	Date      int    `form:"date"`   // '日期',
	Type      string `form:"type"`   // '榜单类型:每日上新mrsx、一周新品yzxp、一周热销yzrx、热销单品rxdp、今日人气jrrq、今日热卖jrrm',
	Stime     int                    // '榜单开始时间',
	Etime     int                    //'榜单结束时间',
	Status    int    `form:"status"` // '状态描述：1：正常，2：关闭，-1：删除',
	Create_id string                 // '创建者Id',
	Update_id string                 //'更新者Id',
	Created   int                    // '创建时间',
	Updated   int                    // '更新时间',
}

func (t VenusTopRankGoodsControl) GetTopRankGoodsControl() *[]VenusTopRankGoodsControl {
	o := orm.NewOrm();
	fmt.Println(t.Date, t.Type)
	r := o.Raw(queryTopRankList, t.Date, t.Type);

	var res []VenusTopRankGoodsControl

	num, err := r.QueryRows(&res)

	if err != nil {
		fmt.Println(err)
		fmt.Println("query info line num:", num)
	}

	fmt.Println(res)

	return &res
}

//
func (t VenusTopRankGoodsControl) GetGoods() {
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw(queryTopRankList, t.Date, t.Type).Values(&maps)
	if err == nil && num > 0 {
		fmt.Println(maps[0]["date"])
	}

	var lists []orm.ParamsList
	num, err = o.Raw(queryTopRankList, t.Date, t.Type).ValuesList(&lists)

	if err == nil && num > 0 {
		fmt.Println(lists[0][0])
	}
}

func (t VenusTopRankGoodsControl) UpdateGoods() int {
	o := orm.NewOrm()
	err1 := o.Begin()
	if err1 != nil {
		err1 = errors.New("事务开启异常！")
	}
	res, err := o.Raw(updateTopRank, t.Update_id, t.Id).Exec()
	if err == nil && res.RowsAffected() > 0 {
		fmt.Println("更新数据影响行数")

		//提交事务
		o.Commit()
		return res.RowsAffected()
	}

	//回滚事务
	o.Rollback()

	return 0
}

