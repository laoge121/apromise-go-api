package base

import (
	"fmt"
)

type person struct {
	Name      string
	Id        int
	Constract struct {
		Phone, city string
	}
}

type student struct {
	person
	Id  int
	sex string
}

type teacher struct {
	person
	Id  int
	sex string
}

func StructT() {
	p := &person{Name: "aaa", Id: 1}
	p.Id = 2
	p.Constract.city = "济南"
	p.Constract.Phone = "1111"
	set1(p)
	set2(p)
	fmt.Println(p)

	t := teacher{person: person{Name: "徐", Id: 1}, sex: "nan", Id: 5}
	s := student{person: person{Name: "李", Id: 2}, sex: "女", Id: 6}
	t.Constract.city = "beij"
	t.Id = 1
	fmt.Println(t)
	fmt.Println(s)
}

func set1(p *person) {
	p.Id = 12
	fmt.Println(p)
}

func set2(p *person) {
	p.Id = 14
	fmt.Println(p)
}
