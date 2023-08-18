package main

import (
	. "fmt"
	_ "fmt"
	log "fmt"
)

func main() {
	// 使用 import alias xxx 时，使用包别名调用包内方法
	log.Println("HelloAlias")
	// 使用 import . xxx 时，可以省略包名
	Println("省略包名")
	// 使用 import _ xxx 时，如果使用包中方法，则编译会报错
	fmt.Println("禁止使用包中方法")
}
