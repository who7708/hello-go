package main

import (
	"fmt"
	"strings"
	"time"
)

type LogProcess struct {
	// channel
	rc chan string
	wc chan string
	// 读取文件的路径
	path string
	// influx data source
	influxDBDsn string
}

// 读取模块
func (lp *LogProcess) ReadFromFile() {
	line := "message"
	lp.rc <- line
}

// 解析模块
func (lp *LogProcess) Process() {
	data := <-lp.rc
	lp.wc <- strings.ToUpper(data)
}

// 写入模块
func (lp *LogProcess) WriteToInfluxDB() {
	fmt.Println(<-lp.wc)
}

func main() {
	lp := &LogProcess{
		// 初始化
		rc:          make(chan string),
		wc:          make(chan string),
		path:        "/tmp/access.log",
		influxDBDsn: "username&password...",
	}

	// 并发执行
	// go (*lp).ReadFromFile()
	go lp.ReadFromFile()
	go lp.Process()
	go lp.WriteToInfluxDB()

	time.Sleep(1 * time.Second)
}
