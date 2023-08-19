package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// 读取接口
type Reader interface {
	Read(rc chan []byte)
}

// 写入接口
type Writer interface {
	Write(wc chan string)
}

type LogProcess struct {
	// channel
	rc chan []byte
	wc chan string
	// // 读取文件的路径
	// path string
	// // influx data source
	// influxDBDsn string
	read  Reader
	write Writer
}

type ReadFromFile struct {
	// 文件路径
	path string
}

type WriteToInfluxDB struct {
	// influx data source
	influxDBDsn string
}

// 实现读取接口
func (r *ReadFromFile) Read(rc chan []byte) {
	// 1. 打开文件
	// 2. 从文件末尾开始逐行读取
	// 3. 写入 Read Channel 中
	f, err := os.Open(r.path)
	if err != nil {
		panic(fmt.Sprintf("open file error: %s", err.Error()))
	}

	f.Seek(0, 2)
	rd := bufio.NewReader(f)

	for {
		line, err := rd.ReadBytes('\n')
		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			fmt.Println("read file error", err)
		}
		rc <- line[:len(line)-1]
	}
}

// 实现写入接口
func (w *WriteToInfluxDB) Write(wc chan string) {
	for data := range wc {
		fmt.Println(data)
	}
}

// // 读取模块
// func (lp *LogProcess) ReadFromFile() {
// }

// 解析模块
func (lp *LogProcess) Process() {
	for data := range lp.rc {
		lp.wc <- strings.ToUpper(string(data))
	}
}

// // 写入模块
// func (lp *LogProcess) WriteToInfluxDB() {
// 	fmt.Println(<-lp.wc)
// }

func main() {
	r := &ReadFromFile{
		path: "./access.log",
	}

	w := &WriteToInfluxDB{
		influxDBDsn: "username&password...",
	}

	lp := &LogProcess{
		// 初始化
		rc: make(chan []byte),
		wc: make(chan string),
		// path:        "/tmp/access.log",
		// influxDBDsn: "username&password...",
		read:  r,
		write: w,
	}

	// 并发执行
	// go (*lp).ReadFromFile()
	go lp.read.Read(lp.rc)
	go lp.Process()
	go lp.write.Write(lp.wc)

	time.Sleep(100 * time.Second)
}
