package main

import (
	"bufio"
	"channel-sample/pipline"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// p := createPipline("small.in", 512, 4)
	// writeToFile(p, "small.out")
	// printFile("small.out")

	// 单机版
	// p := createPipline("large.in", 800_000_000, 8)
	// writeToFile(p, "large.out")
	// // printFile("large.out")

	// 网络版
	// p := createNetworkPipline("small.in", 512, 4)
	// writeToFile(p, "small.out")
	// printFile("small.out")
	p := createNetworkPipline("large.in", 800_000_000, 8)
	writeToFile(p, "large.out")
	printFile("large.out")

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

func createNetworkPipline(filename string, fileSize, chunkCount int) <-chan int {

	chunkSize := fileSize / chunkCount
	pipline.Init()

	sortAddr := []string{}

	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)

		source := pipline.ReaderSource(bufio.NewReader(file), chunkSize)

		addr := ":" + strconv.Itoa(7000+i)
		pipline.NetworkSink(addr, pipline.InMemSort(source))

		sortAddr = append(sortAddr, addr)
	}

	// // 测试
	// return nil

	sortResults := []<-chan int{}
	for _, addr := range sortAddr {
		sortResults = append(sortResults, pipline.NetworkSource(addr))
	}
	return pipline.MergeN(sortResults...)
}
