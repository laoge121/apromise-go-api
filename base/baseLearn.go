package base

import (
	"fmt"
	"strconv"
)

var (
	a int
	b int
	s string
)

type (
	文本 string
)

const (
	C1 = "a"
	C2
	C3 = iota
	C4
)

func ConstCh() {
	fmt.Println(C1)
	fmt.Println(C2)
	fmt.Println(C3)
	fmt.Println(C4)

}

func VarTypeChange() {
	var a int = 65
	b := strconv.Itoa(a)
	a, _ = strconv.Atoi(b)
	fmt.Println(b)
	fmt.Println(a)
}

func VarTest() {
	a = 1
	b = a
	s = "aaaa"
	var data 文本
	data = "中国文字"
	dat := "aaaa"
	fmt.Println(dat)
	fmt.Println(data)
	//var q, r, t, y = 1, 2, 3, 4
	//var q, _, t, y int = 1, 2, 3, 4
	q, r, _, y := 1, 2, 3, 4

	fmt.Println(q, r, y)
	//变量声明
	i := 1.1
	fmt.Println(i)

	fl := 32.1
	ib := 10
	ib = int(fl)

	fmt.Println(ib)
}
