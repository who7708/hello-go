package pkg1

import (
	"fmt"
	"test2/pkg2"
)

func init() {
	fmt.Println("pkg1.init")
}

func Hello() {
	fmt.Println("pkg1.Hello")
	pkg2.Hello()
}
