package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("helloworldserver")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello Server %s</h1>", r.FormValue("name"))
	})

	http.ListenAndServe(":8000", nil)
}
