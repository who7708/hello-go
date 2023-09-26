package main

import (
	"fmt"
	"strings"
)

func stringSample() {
	s := "hello world"

	// 是否包含
	fmt.Println(strings.Contains(s, "hello"), strings.Contains(s, "?"))

	// 索引 base 0
	fmt.Printf("strings.Index(s, \"o\"): %v\n", strings.Index(s, "o"))

	ss := "1#2#345"

	// 切割字符串
	splitStr := strings.Split(ss, "#")
	fmt.Printf("strings.Split(s, \"#\"): %v\n", splitStr)

	// 合并字符串
	fmt.Printf("strings.Join(splitStr, \"#\"): %v\n", strings.Join(splitStr, "#"))

	// 前缀
	fmt.Printf("strings.HasPrefix(\"s\", \"he\"): %v\n", strings.HasPrefix(s, "he"))

	// 后缀
	fmt.Printf("strings.HasSuffix(s, \"d\"): %v\n", strings.HasSuffix(s, "d"))

}
