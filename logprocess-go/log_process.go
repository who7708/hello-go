package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/influxdata/influxdb/client/v2"
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

// 系统状态监控
type SystemInfo struct {
	HandleLine   int     `json:"handleLine"`   // 总处理日志行数
	Tps          float64 `json:"tps"`          // 系统吞出量
	ReadChanLen  int     `json:"readChanLen"`  // read channel 长度
	WriteChanLen int     `json:"writeChanLen"` // write channel 长度
	RunTime      string  `json:"runTime"`      // 运行总时间
	ErrNum       int     `json:"errNum"`       // 错误数
}

const (
	TypeHandleLine = 0
	TypeErrNum     = 1
)

var TypeMonitorChan = make(chan int, 200)

type Monitor struct {
	startTime time.Time
	data      SystemInfo
	tpsSli    []int
}

// 开启监控
func (m *Monitor) start(lp *LogProcess) {

	go func() {
		for n := range TypeMonitorChan {
			switch n {
			case TypeErrNum:
				m.data.ErrNum++
			case TypeHandleLine:
				m.data.HandleLine++
			}
		}
	}()

	ticker := time.NewTimer(time.Second * 5)
	go func() {
		for {
			<-ticker.C
			m.tpsSli = append(m.tpsSli, m.data.HandleLine)
			if len(m.tpsSli) > 2 {
				m.tpsSli = m.tpsSli[1:]
			}
		}
	}()

	http.HandleFunc("/monitor", func(w http.ResponseWriter, r *http.Request) {
		m.data.RunTime = time.Now().Sub(m.startTime).String()
		m.data.ReadChanLen = len(lp.rc)
		m.data.WriteChanLen = len(lp.wc)

		if len(m.tpsSli) > 2 {
			m.data.Tps = float64(m.tpsSli[1]-m.tpsSli[0]) / 5
		}

		ret, _ := json.MarshalIndent(m.data, "", "  ")

		io.WriteString(w, string(ret))
	})

	// listen 方法是阻塞的
	http.ListenAndServe(":9193", nil)
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

		TypeMonitorChan <- TypeHandleLine
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
			TypeMonitorChan <- TypeErrNum
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
			TypeMonitorChan <- TypeErrNum
			log.Println("正则解析失败：", string(data))
			continue
		}

		msg := &Message{}
		t, err := time.ParseInLocation("02/Jan/2006:15:04:05 +0000", ret[4], loc)
		// t, err := time.ParseInLocation("20/Aug/2023:02:30:30 +0000", ret[4], loc)
		if err != nil {
			TypeMonitorChan <- TypeErrNum
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
			TypeMonitorChan <- TypeErrNum
			log.Println("strings.Split fail", ret[6])
			continue
		}
		msg.Method = reqSli[0]

		// url 解析
		u, err := url.Parse(reqSli[1])
		if err != nil {
			TypeMonitorChan <- TypeErrNum
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
	// go run log_process.go -path ./access.log -influxDsn "http://192.168.1.3:8086@admin@admin1234@testdb@s"
	var path, influxDsn string
	var pn, wn int
	flag.StringVar(&path, "path", "./access.log", "读取文件路径")
	flag.StringVar(&influxDsn, "influxDsn", "http://192.168.1.3:8086@admin@admin1234@testdb@s", "InfluxDB数据连接url")
	flag.IntVar(&pn, "pn", 2, "日志处理并发数")
	flag.IntVar(&wn, "wn", 4, "日志写入并发数")
	flag.Parse()

	r := &ReadFromFile{
		path: path,
	}

	w := &WriteToInfluxDB{
		influxDBDsn: influxDsn,
	}

	lp := &LogProcess{
		// 初始化
		rc: make(chan []byte, 200),
		wc: make(chan *Message, 200),
		// path:        "/tmp/access.log",
		// influxDBDsn: "username&password...",
		read:  r,
		write: w,
	}

	// 并发执行
	// go (*lp).ReadFromFile()
	go lp.read.Read(lp.rc)
	for i := 0; i < pn; i++ {
		go lp.Process()
	}
	for i := 0; i < wn; i++ {
		go lp.write.Write(lp.wc)
	}

	m := &Monitor{
		startTime: time.Now(),
		data:      SystemInfo{},
	}

	m.start(lp)
}
