package base

import (
	"fmt"
	"runtime"
	"sync"
	"time"
	//"time"
)

func ThreadT() {

	cha := make(chan int, 10)

	go Gotos(cha)

	fmt.Println(<-cha)

	go func() {
		fmt.Println(">>>>>>>")
	}()

	go fibonacci(cap(cha), cha)

	for val := range cha {
		fmt.Println(val)
	}

}

func Gotos(ch chan int) {
	ch <- 1
	fmt.Println("gooogogogo")
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

//验证go chan 有缓存和没有缓存的区别
func GoBockT() {
	c := make(chan bool)

	go func() {
		fmt.Println("???????go!")
		<-c
	}()
	c <- true
	fmt.Println("ces")

	d := make(chan bool, 1)
	go func() {
		fmt.Println("???????go2!")
		d <- true
	}()
	<-d
}

func GoThread() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go print(i, c)
	}
	for i := 0; i < 10; i++ {
		<-c
	}
}

func print(i int, c chan bool) {
	fmt.Println("aaaaaa", i)
	c <- true
}

func prints(i int, wg *sync.WaitGroup) {
	fmt.Println("bbbbb", i)
	wg.Done()

}

func GoSys() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go prints(i, &wg)
	}
	wg.Wait()
}

func SelectT() {
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					fmt.Println(">>>>>>c1")
					o <- true
					break
				}
				fmt.Println(">>>1", v)
			case v, ok := <-c2:
				if !ok {
					fmt.Println(">>>>>>c2")
					o <- true
					break
				}
				fmt.Println(">>>2", v)
			}
		}
	}()
	c1 <- 1
	c2 <- "aaa"
	c1 <- 2
	c2 <- "bbb"
	close(c1)
	close(c2)
	<-o
}

func SelectT1() {
	c1 := make(chan int)

	go func() {
		for v := range c1 {
			fmt.Println(v)
		}
	}()

	for {
		select {
		case c1 <- 0:
		case c1 <- 1:
		}
	}

}

func SelectTimeOut() {

	c := make(chan int)
	o := make(chan bool)
	select {
	case v := <-c:
		fmt.Println(v)
	case <-time.After(3 * time.Second):
		fmt.Println("aaa")
		o <- true
		break
	}
	<-o
}

//单向chan 的声明使用
func ChanType() {

	m := make(chan int)

	//r := <-chan int(m) //单向读chan
	w := chan<- int(m) //单向写chan

	w <- 1
	//a := <-r
	<-m
}

/**
相当于java的lock 不是 比较保利
**/
func LockT() {
	var lock1 sync.Mutex
	lock1.Lock()

	defer lock1.Unlock()

	var lock2 sync.RWMutex

	lock2.RLock()
	defer lock2.RUnlock()

	lock2.Lock()

	defer lock2.Unlock()

	//让出时间片段
	runtime.Gosched()
}

func oncep() {
	fmt.Println("oncep")
}

//全局唯一性的操作
func OnceT() {
	var once sync.Once
	once.Do(oncep)

}
