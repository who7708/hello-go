package main

import (
	"fmt"
	"net/http"
	"time"
	"ws-go/impl"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// // 写内容
	// w.Write([]byte("hello"))
	fmt.Println("Upgrade 开始建立连接")

	var (
		wsConn *websocket.Conn
		err    error
		data   []byte
		conn   *impl.Connection
	)
	// Upgrade: websocket
	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		fmt.Println("Upgrade 异常")
		return
	}

	if conn, err = impl.InitConnection(wsConn); err != nil {
		goto ERR
	}

	go func() {
		var (
			err error
		)
		for {
			if err = conn.WriteMessage([]byte("heartbeat")); err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}

	}

	// //  websocket.wsConn
	// for {
	// 	//  Text, Binary
	// 	// 读取消息
	// 	if _, data, err = wsConn.ReadMessage(); err != nil {
	// 		fmt.Println("读取异常")
	// 		goto ERR
	// 	}
	// 	// 回写消息
	// 	if err = wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
	// 		fmt.Println("写入消息")
	// 		goto ERR
	// 	}
	// }

ERR:
	// 关闭连接的操作
	// wsConn.Close()

}

func main() {
	http.HandleFunc("/ws", wsHandler)

	http.ListenAndServe(":8000", nil)
}
