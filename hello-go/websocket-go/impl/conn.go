package impl

import (
	"errors"
	"sync"

	"github.com/gorilla/websocket"
)

type Connection struct {
	wsConn    *websocket.Conn
	inChan    chan []byte
	outChan   chan []byte
	closeChan chan byte
	mutex     sync.Mutex
	isClosed  bool
}

// 初始化长连接
func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConn: wsConn,
		// 容量 1000
		inChan:    make(chan []byte, 1000),
		outChan:   make(chan []byte, 1000),
		closeChan: make(chan byte, 1),
	}

	// 启动读协程
	go conn.readLoop()

	// 启动写协程
	go conn.writeLoop()

	return
}

// API
// 读取消息
func (conn *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
		err = errors.New("连接已关闭")
	}
	return
}

// 写入消息
func (conn *Connection) WriteMessage(data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("连接已关闭")
	}
	return
}

// 关闭连接,要线程安全
func (conn *Connection) Close() {
	// Close 方法是线程安全
	conn.wsConn.Close()

	conn.mutex.Lock()
	if !conn.isClosed {
		// 关闭channel
		// 需要保证此 close 只执行一次
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

// 内部实现
func (conn *Connection) readLoop() {
	var (
		data []byte
		err  error
	)
	// 循环读取消息,并写入chnnel中
	for {
		if _, data, err = conn.wsConn.ReadMessage(); err != nil {
			goto ERR
		}
		// 阻塞在这里,等待inChan 有空闲位置
		select {
		case conn.inChan <- data:
		case <-conn.closeChan:
			// closeChan 关闭的时候
			goto ERR
		}

	}
ERR:
	conn.Close()
}

func (conn *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)
	// 循环写入消息,并写入chnnel中
	for {
		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			// closeChan 关闭的时候
			goto ERR
		}

		if err = conn.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}
