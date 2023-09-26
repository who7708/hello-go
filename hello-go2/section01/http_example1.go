package section01

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("started %v %v", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("executing index handler")
	fmt.Fprintf(w, "welcome")
}

func about(w http.ResponseWriter, r *http.Request) {
	log.Println("executing about handler")
	fmt.Fprintf(w, "go middleware")
}

func iconHandler(w http.ResponseWriter, r *http.Request) {

}

func Run() {
	http.HandleFunc("/favicon.ico", iconHandler)

	indexHandler := http.HandlerFunc(index)
	aboutHandler := http.HandlerFunc(about)

	http.Handle("/", loggingHandler(indexHandler))
	http.Handle("/about", loggingHandler(aboutHandler))

	server := &http.Server{
		Addr: ":8001",
	}

	log.Println("Listening...")
	server.ListenAndServe()
}
