package main

import (
	"fmt"
	"time"
	"unsafe"
)

func TestInt8() {
	var i int8 = 1
	fmt.Println(unsafe.Sizeof(i))
}

func TestInt16() {
	var i int16 = 1
	fmt.Println(unsafe.Sizeof(i))
}

func TestInt32() {
	var i int32 = 1
	fmt.Println(unsafe.Sizeof(i))
}

func TestInt64() {
	var i int64 = 1
	fmt.Println(unsafe.Sizeof(i))
}

func TestFloat32() {
	var i float32 = 1
	fmt.Println(unsafe.Sizeof(i))
}

func TestFloat64() {
	var i float64 = 1
	fmt.Println(unsafe.Sizeof(i))
}

func TestConst() {
	const str1 string = "我是字符串常量"
	fmt.Println(str1)
	const str2 = "我是隐式字符串常量"
	fmt.Println(str2)

	const (
		cat1 string = "猫"
		cat2        = "狗"
	)
	fmt.Println(cat1)
	fmt.Println(cat2)

	const apple, banana string = "苹果", "香蕉"
	fmt.Println(apple)
	fmt.Println(banana)

	const a, b, c = 123, "香蕉", "hello"
	fmt.Println(a, b, c)

	// 每个中文占3个字节，输出 6
	const lenB = len(b)
	fmt.Println(lenB)

	// 每个字母占1个字节，输出 5
	const lenC = len(c)
	fmt.Println(lenC)

}

func TestIota() {
	const a = iota

	const (
		b = iota
		c = iota
	)
	// 0 0 1
	fmt.Println(a, b, c)

	const (
		d, e = iota, iota * 3
		f, g
		h, j
	)
	// 0 0 1 3 2 6
	fmt.Println(d, e, f, g, h, j)
}

func TestSwitch() {
	var a interface{}
	a = 32
	switch a.(type) {
	case int:
		fmt.Println("整数类型")
	case string:
		fmt.Println("字符串类型")
	default:
		fmt.Println("其他类型")
	}
}

func main() {
	TestInt8()
	TestInt16()
	TestInt32()
	TestInt64()
	TestFloat32()
	TestFloat64()

	fmt.Println("============================")
	i := byte(1)
	fmt.Println(i)

	i1 := int32(1)
	fmt.Println(i1)

	fmt.Println("============================")

	var b int64 = 123
	fmt.Println(unsafe.Sizeof(b))
	// 强制类型转换
	a := int32(b)
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(a)

	fmt.Println("============================")
	a1, b1, c1 := 11, 22, 33
	fmt.Println(a1, b1, c1)

	fmt.Println("============================")
	a2 := 123
	fmt.Println(a2)

	fmt.Println("============================")
	TestConst()

	fmt.Println("============================")
	TestIota()

	fmt.Println("============================")
	TestSwitch()

	fmt.Println("============================")
	TestFor3()

	fmt.Println("============================")
	TestGoto()

}

func TestGoto() {
	// 下面代码会形成死循环
	// 	goto One
	// Two:
	// 	fmt.Println("world")
	// One:
	// 	fmt.Println("hello")
	// 	goto Two
}

func TestFor1() {
	// panic("unimplemented")
	// 类似 while true
	for {
		fmt.Println("Hello")
		time.Sleep(1 * time.Second)
	}
}

func TestFor2() {
	// panic("unimplemented")
	for i := 0; i < 10; i++ {
		fmt.Println("Hello")
		time.Sleep(1 * time.Second)
	}
}

func TestFor3() {
	a := []string{"apple", "banana"}
	for idx, value := range a {
		fmt.Println(idx)
		fmt.Println(value)
	}
	// for _, value := range a {
	// 	fmt.Println(value)
	// }
}
