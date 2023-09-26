package main

import (
	"fmt"
	"time"
)

func main() {
	// testGoRoutine()
	testGoRoutine2()
}

// goroutine
func testGoRoutine() {
	for i := 0; i < 5; i++ {
		// go starts a goroutine
		go printHelloWorld(i)
	}
	time.Sleep(1 * time.Second)
}

func printHelloWorld(i int) {
	for {
		fmt.Printf("Hello World goroutine %d \n", i)
	}
}

// 管道
func testGoRoutine2() {
	ch := make(chan string)
	for i := 0; i < 5; i++ {
		// go starts a goroutine
		go printHelloWorld2(i, ch)
	}

	for {
		msg := <-ch
		fmt.Println(msg)
	}
}

func printHelloWorld2(i int, ch chan string) {
	for {
		ch <- fmt.Sprintf("Hello World goroutine %d \n", i)
	}
}
