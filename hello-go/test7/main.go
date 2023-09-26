package main

import (
	"fmt"
	"sync"
	"time"
)

var i int32 = 0

// ++ 操作, 非线程安全
func doAdd() {
	// i++ 非原子操作,所以在多线程下会不正确
	i++
}

// ++ 操作
func doAdd2(c chan int32) {
	// i++ 非原子操作,所以在多线程下会不正确
	i++
	c <- i
}

// 耗时打印
func timeCost(start time.Time) {
	cost := time.Since(start)
	fmt.Printf("cost: %v\n", cost)
}

func test1() {

	begin := time.Now()
	for i := 0; i < 50000; i++ {
		go doAdd()
	}
	cost := time.Since(begin) //.Milliseconds()
	fmt.Printf("i: %v , cost %v\n", i, cost)
	fmt.Printf("i: %+v , cost %+v\n", i, cost)
	fmt.Printf("i: %#v , cost %#v\n", i, cost)
	// 上面正常输出
	// i: int , cost int64
	fmt.Printf("i: %T , cost %T\n", i, cost)
	// i: % , cost %
	fmt.Printf("i: %% , cost %%\n")
}

// defer 方式打印耗时
func test2() {

	defer timeCost(time.Now())
	for i := 0; i < 50000; i++ {
		go doAdd()
	}
	fmt.Printf("i: %v\n", i)
}

var (
	rwlock sync.RWMutex
)

func main() {

	ch := make(chan int32)
	defer timeCost(time.Now())
	for i := 0; i < 50000; i++ {
		// 还是一样无法保证线程安全,为啥呢?
		rwlock.TryLock()
		go doAdd2(ch)
		rwlock.Unlock()
	}
	fmt.Printf("i: %v\n", <-ch)

}
