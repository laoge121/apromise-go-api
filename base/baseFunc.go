package base

import (
	"fmt"
)

/**
 不定长变参 必须是最后一个参数 是值传递
**/
func FuncT(a ...int) {

	a[0] = 1
	a[1] = 3
	fmt.Println(a)

}

/**
匿名函数
**/
func Change(x int) func(int) int {
	x = 9
	return func(y int) int {
		return x * y
	}
}

func DeferT() {
	fmt.Println("a")
	//defer fmt.Println("b")
	fmt.Println("c")

	//for i := 0; i < 3; i++ {
	//	defer fmt.Println(i)
	//}
	/**
		for i := 1; i < 10; i++ {
			defer func() {
				fmt.Println(i)
			}()
		}
	**/

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("恢复结果！")
		}
	}()
	panic("this func exception!")
}
