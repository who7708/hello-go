package main

import (
	"fmt"
	"test2/pkg1"
)

func init() {
	fmt.Println("main.init")
}

func main() {
	fmt.Println("main")
	pkg1.Hello()
}
