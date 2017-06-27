package base

import (
	"fmt"
)

type strB string

type IntT int

type A struct {
	Name string
}

type B struct {
	Name string
}

func (i *IntT) Increment() {
LOOP:
	for {
		*i += IntT(a)
		if int(*i) == 100 {
			break LOOP
		}
	}
}

func (s *strB) StrBT() {
	fmt.Println(s)
}

func (a A) Println() {
	fmt.Println(a.Name + "-a")
}

func (b B) Println() {
	fmt.Println(b.Name + "-b")
}

func (a *A) Println1() {
	a.Name = "aaa"
	fmt.Println(a.Name + "-a")
}

func (b *B) Println2() {
	b.Name = "bbb"
	fmt.Println(b.Name + "-b")
}
