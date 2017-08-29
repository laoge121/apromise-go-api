package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

var (
	queryTopRankList = "SELECT id,date,type,stime,etime,status,create_id,update_id,created,updated from t_venus_top_rank_goods_control where date=? and type =?  and status=2 limit 10 "
)

type VenusTopRankGoodsControl struct {
	Id        int    // '自增主键',
	Date      int    // '日期',
	Type      string // '榜单类型:每日上新mrsx、一周新品yzxp、一周热销yzrx、热销单品rxdp、今日人气jrrq、今日热卖jrrm',
	Stime     int    // '榜单开始时间',
	Etime     int    //'榜单结束时间',
	Status    int    // '状态描述：1：正常，2：关闭，-1：删除',
	Create_id string // '创建者Id',
	Update_id string //'更新者Id',
	Created   int    // '创建时间',
	Updated   int    // '更新时间',
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
