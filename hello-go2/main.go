package main

import (
	"fmt"
	"learn/graph"
	"learn/section05"
	"log"
	"net/http"
)

type dollars float32

func main() {
	// section01.Section01()
	// section01.HttpGetTest()
	// section01.HttpPostTest()
	// section01.HttpWebTest()
	// run()
	// section01.Run()
	// section02.Run()

	// section03.Run()
	// section04.Run()

	// a := section03.App{}
	// a.Initialize()
	// a.Run(":8001")

	// testPath()

	// base.Run()
	// base.Run1()

	graph.GraphRun()
}

func testPath() {
	section05.SetConfig("test.cfg")
	section05.SetHomeDir("")
}

func run() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8001", mux))
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%v : %v \n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
