package main

import (
	"bufio"
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
	"io"
	"log"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 读取接口
type Reader interface {
	Read(rc chan []byte)
}

// 写入接口
type Writer interface {
	Write(wc chan *Message)
}

type LogProcess struct {
	// channel
	rc chan []byte
	wc chan *Message
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

// 解析数据格式
type Message struct {
	TimeLocal                    time.Time
	ByteSent                     int32
	Path, Method, Scheme, Status string
	UpstreamTime, RequestTime    float64
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
func (w *WriteToInfluxDB) Write(wc chan *Message) {
	// url@用户名@密码@数据库@精度
	// "http://192.168.1.3:8086@admin@admin1234@testdb@s"
	infSli := strings.Split(w.influxDBDsn, "@")

	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     infSli[0],
		Username: infSli[1],
		Password: infSli[2],
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	for data := range wc {
		// Create a new point batch
		bp, err := client.NewBatchPoints(client.BatchPointsConfig{
			Database:  infSli[3],
			Precision: infSli[4],
		})
		if err != nil {
			log.Fatal(err)
		}

		// Create a point and add to batch
		// tags: Path, Method, Scheme, Status
		// fields: UpstreamTime, RequestTime, ByteSent
		// Time: TimeLocal
		tags := map[string]string{
			"Path":   data.Path,
			"Method": data.Method,
			"Scheme": data.Scheme,
			"Status": data.Status,
		}
		fields := map[string]interface{}{
			"UpstreamTime": data.UpstreamTime,
			"RequestTime":  data.RequestTime,
			"ByteSent":     data.ByteSent,
		}

		pt, err := client.NewPoint("nginx_log", tags, fields, data.TimeLocal)
		if err != nil {
			log.Fatal(err)
		}
		bp.AddPoint(pt)

		// Write the batch
		if err := c.Write(bp); err != nil {
			log.Fatal(err)
		}

		// // Close client resources
		// if err := c.Close(); err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Println(data)
		log.Println("write success!")
	}
}

// // 读取模块
// func (lp *LogProcess) ReadFromFile() {
// }

// 解析模块
func (lp *LogProcess) Process() {

	/**
	172.0.0.12 - - [04/Mar/2018:13:49:52 +0000] http "GET /foo?query=t HTTP/1.0" 200 2133 "-" "KeepAliveClient" "-" 1.005 1.854
	*/

	r := regexp.MustCompile(`([\d\.]+)\s+([^ \[]+)\s+([^ \[]+)\s+\[([^\]]+)\]\s+([a-z]+)\s+\"([^"]+)\"\s+(\d{3})\s+(\d+)\s+\"([^"]+)\"\s+\"(.*?)\"\s+\"([\d\.-]+)\"\s+([\d\.-]+)\s+([\d\.-]+)`)

	loc, _ := time.LoadLocation("Asia/Shanghai")

	for data := range lp.rc {
		ret := r.FindStringSubmatch(string(data))
		if len(ret) != 14 {
			log.Println("正则解析失败：", string(data))
			continue
		}

		msg := &Message{}
		t, err := time.ParseInLocation("02/Jan/2006:15:04:05 +0000", ret[4], loc)
		// t, err := time.ParseInLocation("20/Aug/2023:02:30:30 +0000", ret[4], loc)
		if err != nil {
			log.Println("时间解析出错：", err.Error(), ret[4])
			continue
		}
		msg.TimeLocal = t

		// bytesent
		byteSent, _ := strconv.Atoi((ret[8]))
		msg.ByteSent = int32(byteSent)

		// Path, Method, Scheme, Status
		// GET /foo?query=t HTTP/1.0
		reqSli := strings.Split(ret[6], " ")
		if len(reqSli) != 3 {
			log.Println("strings.Split fail", ret[6])
			continue
		}
		msg.Method = reqSli[0]

		// url 解析
		u, err := url.Parse(reqSli[1])
		if err != nil {
			log.Println("url parse fail:", err)
			continue
		}
		msg.Path = u.Path

		msg.Scheme = ret[5]
		msg.Status = ret[7]

		upstreamTime, _ := strconv.ParseFloat(ret[12], 64)
		RequestTime, _ := strconv.ParseFloat(ret[13], 64)

		msg.UpstreamTime = upstreamTime
		msg.RequestTime = RequestTime

		// UpstreamTime, RequestTime    float64

		lp.wc <- msg
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
		influxDBDsn: "http://192.168.1.3:8086@admin@admin1234@testdb@s",
	}

	lp := &LogProcess{
		// 初始化
		rc: make(chan []byte),
		wc: make(chan *Message),
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
