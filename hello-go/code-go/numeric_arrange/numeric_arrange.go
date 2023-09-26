package main

import (
	"fmt"
)

func main() {
	totalCount := 0
	/*以下为三重循环*/
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			for k := 1; k < 5; k++ {
				/*确保 i 、j 、k 三位互不相同*/
				if i != k && i != j && j != k {
					totalCount++
					fmt.Println("第", totalCount, "方案", "i =", i, "j =", j, "k =", k)
				}
			}
		}
	}
	fmt.Println("共", totalCount, "种方案")
}
