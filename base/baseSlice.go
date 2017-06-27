package base

import (
	"fmt"
)

func SliceT() {
	arr := [...]int{1, 2, 3, 4, 52, 56, 3, 7, 8, 9, 0}
	var sli1 []int

	sli1 = arr[:len(arr)]

	sli2 := arr[:5]
	fmt.Println(sli1)
	fmt.Println(sli2)

	sli3 := make([]int, 10, 20)

	fmt.Println(len(sli3), cap(sli3))

	s1 := make([]int, 3, 6)
	fmt.Println(s1, &s1)
	s1 = append(s1, 1, 2, 3)
	fmt.Println(s1, &s1)
	s1 = append(s1, 1, 2, 3)
	fmt.Println(s1, &s1)

	a := [...]int{1, 2, 3, 4, 5}
	sa1 := a[2:5]
	sa2 := a[1:3]
	fmt.Println(sa1, sa2)
	sa1[0] = 9
	fmt.Println(sa1, sa2)

	a1 := [...]int{1, 2, 3, 4, 5}
	sa11 := a1[2:5]
	sa21 := a1[1:3]
	fmt.Println(sa11, sa21)
	sa21 = append(sa21, 1, 1, 1, 1, 2, 3, 3)
	sa11[0] = 9
	fmt.Println(sa11, sa21)

	//copy
	c1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	c2 := []int{11, 22, 33}
	//copy(c1, c2)
	//fmt.Println(c1)

	copy(c2, c1)
	fmt.Println(c2)

	//copy(c1, c2)
	//fmt.Println(c1)
}
