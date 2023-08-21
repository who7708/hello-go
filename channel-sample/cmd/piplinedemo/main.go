package main

import (
	"bufio"
	"channel-sample/pipline"
	"fmt"
	"os"
)

func main() {
	// mergeDemo()
	readerSourceDemo()
}

func readerSourceDemo() {
	const filename = "small.in"
	const n = 100_000_000
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p := pipline.RandomSource(n)

	writer := bufio.NewWriter(file)
	pipline.WriteSink(writer, p)
	// 将最后buff中的数据flush掉
	writer.Flush()

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	p = pipline.ReaderSource(bufio.NewReader(file))

	// 打印前100个
	count := 0
	for v := range p {
		count++
		if count <= 100 {
			fmt.Println(v)
			continue
		}
		break
	}
}

func mergeDemo() {
	// 内部排序
	// p := pipline.InMemSort(
	// 	pipline.ArraySource(3, 2, 6, 7, 4),
	// )

	// 归并排序
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
