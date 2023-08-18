package main

import (
	"fmt"
	"net/http"

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
		conn *websocket.Conn
		err  error
		data []byte
	)
	// Upgrade: websocket
	if conn, err = upgrader.Upgrade(w, r, nil); err != nil {
		fmt.Println("Upgrade 异常")
		return
	}

	//  websocket.conn
	for {
		//  Text, Binary
		// 读取消息
		if _, data, err = conn.ReadMessage(); err != nil {
			fmt.Println("读取异常")
			goto ERR
		}
		// 回写消息
		if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
			fmt.Println("写入消息")
			goto ERR
		}
	}

ERR:
	conn.Close()

}

func main() {
	http.HandleFunc("/ws", wsHandler)

	http.ListenAndServe(":8000", nil)
}
