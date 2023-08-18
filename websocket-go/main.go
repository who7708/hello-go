package main

import "net/http"

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// 写内容
	w.Write([]byte("hello"))
}

func main() {
	http.HandleFunc("/ws", wsHandler)

	http.ListenAndServe(":8000", nil)
}
