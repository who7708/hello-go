package main

import (
	"fmt"
	"time"
)

var i int32 = 0

// ++ 操作
func doAdd() {
	// i++ 非原子操作,所以在多线程下会不正确
	i++
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

func main() {
	defer timeCost(time.Now())
	for i := 0; i < 50000; i++ {
		go doAdd()
	}
	fmt.Printf("i: %v\n", i)
}
