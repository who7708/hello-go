package main

import (
	"bufio"
	"channel-sample/pipline"
	"fmt"
	"os"
)

func main() {
	// p := createPipline("small.in", 100_000_000, 8)
	// writeToFile(p, "small.out")
	// printFile("small.out")
	p := createPipline("large.in", 800_000_000, 8)
	writeToFile(p, "large.out")
	// printFile("large.out")
}

func writeToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipline.WriteSink(writer, p)
}

func printFile(filename string) {

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipline.ReaderSource(file, -1)

	count := 0
	for v := range p {
		count++
		if count >= 100 {
			break
		}
		fmt.Println(v)
	}

}

func createPipline(filename string, fileSize, chunkCount int) <-chan int {

	chunkSize := fileSize / chunkCount
	pipline.Init()

	sortResults := []<-chan int{}

	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)

		source := pipline.ReaderSource(bufio.NewReader(file), chunkSize)

		sortResults = append(sortResults, pipline.InMemSort(source))

	}
	return pipline.MergeN(sortResults...)
}
