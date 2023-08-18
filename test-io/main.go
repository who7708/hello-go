package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ReadFrom(reader io.Reader, num int32) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

// 字符串读取
func sampleReadFromString() {
	data, _ := ReadFrom(strings.NewReader("from string"), 5)
	fmt.Println(string(data))
}

// 标准输入读取
func sampleReadFromStdin() {
	fmt.Println("输入字符:")
	data, _ := ReadFrom(os.Stdin, 2)
	fmt.Println(string(data))
}

// 文件中读取
func sampleReadFromFile() {
	file, _ := os.Open("main.go")
	defer file.Close()

	data, _ := ReadFrom(file, 88)
	fmt.Printf("data: %v\n", string(data))
}

// bufio
func sampleReadFromBufio() {
	strReader := strings.NewReader("hello world")
	bufReader := bufio.NewReader(strReader)

	data, _ := bufReader.Peek(5)

	// buffer 中只读不取
	// data: hello , size 11
	fmt.Printf("data: %v , size %v\n", string(data), bufReader.Buffered())

	str, _ := bufReader.ReadString(' ')
	// buffer 中读完就取出来
	// data: hello  , size 5
	fmt.Printf("data: %v , size %v\n", str, bufReader.Buffered())

	// 写入标准输出
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "hello ")
	fmt.Fprint(w, "world")
	// hello world
	w.Flush()
}

// 统计代码行
func codeCounter() {
	if len(os.Args) < 2 {
		return
	}
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		// fmt.Println(err)
		log.Fatal(err)
		return
	}
	defer file.Close()

	bufReader := bufio.NewReader(file)

	var line int32
	for {
		_, isPrefix, err := bufReader.ReadLine()
		if err != nil {
			break
		}

		// 如果文件行超长,则isPrefix为true
		if !isPrefix {
			line++
		}
	}
	fmt.Printf("line: %v\n", line)
	// log.Fatal(line)
}

// 读取文件头内容
func readFileHeader() {
	file, _ := os.Open("user.bmp")
	defer file.Close()

	// bmp 文件头
	var headA, headB byte
	binary.Read(file, binary.LittleEndian, &headA)
	binary.Read(file, binary.LittleEndian, &headB)

	// 文件大小
	var size uint32
	binary.Read(file, binary.LittleEndian, &size)

	// 保留字节
	var reservedA, reservedB byte
	binary.Read(file, binary.LittleEndian, &reservedA)
	binary.Read(file, binary.LittleEndian, &reservedB)

	// 偏移字节,文件内容从哪个地方开始
	var offBits uint32
	binary.Read(file, binary.LittleEndian, &offBits)

	fmt.Printf("headA: %c\n", headA)
	fmt.Printf("headB: %c\n", headB)

	fmt.Printf("size: %v\n", size)

	fmt.Printf("reservedA: %v\n", reservedA)
	fmt.Printf("reservedB: %v\n", reservedB)

	fmt.Printf("offBits: %v\n", offBits)

}

type BitmapInfoHeader struct {
	Size           uint32
	Width          int32
	Height         int32
	Places         uint16
	BitCount       uint16
	Compression    uint32
	SizeImage      uint32
	XperlsPerMeter int32
	YperlsPerMeter int32
	ClsrUsed       uint32
	ClrImportant   uint32
}

// 使用结构体进行读取
func readFileHeader2() {
	infoHeader := new(BitmapInfoHeader)
	file, _ := os.Open("user.bmp")
	defer file.Close()
	binary.Read(file, binary.LittleEndian, infoHeader)

	fmt.Printf("infoHeader: %v\n", infoHeader)
}

func main() {
	// sampleReadFromString()
	// sampleReadFromStdin()
	// sampleReadFromFile()
	// sampleReadFromBufio()
	// codeCounter()
	// readFileHeader()
	readFileHeader2()
}
