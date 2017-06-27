package base

import (
	"fmt"
)

func PointT() {

	a := 1

	var P *int
	P = &a
	fmt.Println(*P)
}

func IfElseT() {
	a := 10
	if a := 1; a > 2 || a <= 1 {
		fmt.Println("aaa", a)
	}
	fmt.Println("bbb", a)
}

func ForT() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func SwitchT() {
	//a := 1
	switch a := 1; {
	case a >= 0:
		fmt.Println(">>>>>0")
		fallthrough
	case a >= 1:
		fmt.Println(">>>>>>.1")
		fallthrough
	default:
		fmt.Println(">>>>>>defautl")
	}
	switch b := 1; {
	case b >= 0:
		fmt.Println(">>>>>>>>>>0")
		fallthrough
	case b >= 1:
		fmt.Println(">>>>>>.1")
		fallthrough
	default:
		fmt.Println(">>>>>>defautl")
	}
}

func GotoBreakContinueT() {

}
