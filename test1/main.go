package main

import (
	"fmt"
	"reflect"
	"time"
)

// 常量定义
const NAME string = "hello"

// 全局变量声明与赋值
var a string = "hello"

// 一般类型声明
type helloInt int

// 结构声明
type Hello struct{}

// 接口声明
type IHello interface{}

// 自定义函数
func sayHello() {
	fmt.Println("HelloWorld")
	fmt.Println(time.Now())

	var var1 helloInt = 123
	fmt.Printf("reflect.TypeOf(var1): %v\n", reflect.TypeOf(var1))

	var2 := 123
	fmt.Printf("reflect.TypeOf(var2): %v\n", reflect.TypeOf(var2))
}

// main
func main() {
	sayHello()
	fmt.Println("Main Hello World")
}
