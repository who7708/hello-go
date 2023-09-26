package main

import (
	"flag"
	"fmt"
	"os"
)

// 命令行解析
// go run *.go -method hello -value 123
func readArg() {
	// 命令行参数
	s := os.Args
	methodPtr := flag.String("method", "default", "method of sample")
	valuePtr := flag.Int("value", -1, "value of sample")
	flag.Parse()

	// [C:\Users\chmy\AppData\Local\Temp\go-build985525039\b001\exe\arg_sample.exe -method hello -value 123]
	fmt.Println(s)
	// methodPtr: 0xc00000a028
	// valuePtr: 0xc00000a030
	// * 运算符访问指针指向的值
	// 打印指针地址
	fmt.Println(methodPtr, valuePtr)
	// 打印值
	fmt.Println(*methodPtr, *valuePtr)
	fmt.Printf("methodPtr: %v\n", methodPtr)
	fmt.Printf("methodPtr: %v\n", *methodPtr)
	fmt.Printf("valuePtr: %v\n", valuePtr)
	fmt.Printf("valuePtr: %v\n", *valuePtr)
}

// 命令行解析
// go run *.go -method hello -value 123
func readArg2() {
	var method string
	var value int
	// 命令行参数
	flag.StringVar(&method, "method", "default", "method of sample")
	flag.IntVar(&value, "value", -1, "value of sample")
	flag.Parse()

	// 打印指针
	fmt.Println(&method, &value)
	// 打印值
	fmt.Println(method, value)

	fmt.Printf("methodPtr: %v\n", method)
	fmt.Printf("valuePtr: %v\n", value)
}
