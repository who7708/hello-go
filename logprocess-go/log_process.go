package main

import (
	"fmt"
	"strings"
	"time"
)

// 读取接口
type Reader interface {
	Read(rc chan string)
}

// 写入接口
type Writer interface {
	Write(wc chan string)
}

type LogProcess struct {
	// channel
	rc chan string
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
func (r *ReadFromFile) Read(rc chan string) {
	line := "message"
	rc <- line
}

// 实现写入接口
func (w *WriteToInfluxDB) Write(wc chan string) {
	fmt.Println(<-wc)
}

// // 读取模块
// func (lp *LogProcess) ReadFromFile() {
// }

// 解析模块
func (lp *LogProcess) Process() {
	data := <-lp.rc
	lp.wc <- strings.ToUpper(data)
}

// // 写入模块
// func (lp *LogProcess) WriteToInfluxDB() {
// 	fmt.Println(<-lp.wc)
// }

func main() {
	r := &ReadFromFile{
		path: "/tmp/access.log",
	}

	w := &WriteToInfluxDB{
		influxDBDsn: "username&password...",
	}

	lp := &LogProcess{
		// 初始化
		rc: make(chan string),
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

	time.Sleep(1 * time.Second)
}
