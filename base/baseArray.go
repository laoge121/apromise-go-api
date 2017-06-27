package base

import (
	"fmt"
)

func ArratT() {
	var a [2]int
	a[0] = 1
	a[1] = 2
	a1 := [2]int{1: 9}
	var b [2]int
	b[1] = 2
	b[0] = 1
	b1 := [...]int{19: 1}

	var P *[20]int

	P = &b1

	p := new([20]int)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(P)
	fmt.Println(a == b)
	fmt.Println(p)

	d2 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(d2)
}

func PxMp() {
	a := [...]int{4, 2, 1, 3, 2, 6, 7, 8, 9}
	fmt.Println(a)
	num := len(a)

	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if a[i] < a[j] {
				tmp := a[i]
				a[i] = a[j]
				a[j] = tmp
			}
		}
	}
	fmt.Println(a)
}
