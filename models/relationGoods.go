package models

type RelationGoods struct {
	id       int
	goods_id int
	title    string
	stime    float32
}

type common interface{
	test()
	aa()
}
