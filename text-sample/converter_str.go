package main

import (
	"fmt"
	"strconv"
)

func converterStr() {
	fmt.Printf("strconv.Itoa(10): %v\n", strconv.Itoa(10))

	// 711 <nil>
	fmt.Println(strconv.Atoi("711"))

	// false <nil>
	fmt.Println(strconv.ParseBool("false"))

	// 3.1415926 <nil>
	fmt.Println(strconv.ParseFloat("3.1415926", 64))

	// strconv.FormatBool(true): true
	fmt.Printf("strconv.FormatBool(true): %v\n", strconv.FormatBool(true))

	// 转换成4进制 strconv.FormatInt(123, 4): 1323
	fmt.Printf("strconv.FormatInt(123, 4): %v\n", strconv.FormatInt(123, 4))

}
