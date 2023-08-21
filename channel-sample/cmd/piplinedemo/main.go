package main

import (
	"channel-sample/pipline"
	"fmt"
)

func main() {

	// p := pipline.InMemSort(
	// 	pipline.ArraySource(3, 2, 6, 7, 4),
	// )
	p := pipline.Merge(
		pipline.InMemSort(
			pipline.ArraySource(3, 2, 6, 7, 4),
		),
		pipline.InMemSort(
			pipline.ArraySource(7, 4, 0, 3, 2, 13, 8),
		))
	// for {
	// 	if num, ok := <-p; ok {
	// 		fmt.Println(num)
	// 	} else {
	// 		break
	// 	}
	// }

	for v := range p {
		fmt.Println(v)
	}
}
