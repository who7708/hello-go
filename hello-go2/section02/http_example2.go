package section02

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/justinas/alice"
	"log"
	"net/http"
	"os"
	"time"
)

func middlewareFirst(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("started %v %v", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}

func middlewareSecond(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("MiddlewareSecond - before Handler")
		if r.URL.Path == "/message" {
			if r.URL.Query().Get("password") == "123" {
				log.Println("Authorized to system...")
				next.ServeHTTP(w, r)
			} else {
				log.Println("Failed to authorize to the system")
				return
			}
		} else {
			next.ServeHTTP(w, r)
		}
		log.Println("MiddlewareScond - after Handler")
	})
}

func loggingHandler(next http.Handler) http.Handler {

	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}

	return handlers.LoggingHandler(logFile, next)
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("executing index handler")
	fmt.Fprintf(w, "welcome")
}

func message(w http.ResponseWriter, r *http.Request) {
	log.Println("Executed message Handler")
	fmt.Fprintf(w, "message...")
}

func iconHandler(w http.ResponseWriter, r *http.Request) {

}

func Run() {
	http.HandleFunc("/favicon.ico", iconHandler)

	indexHandler := http.HandlerFunc(index)
	messageHandler := http.HandlerFunc(message)

	commonHandlers := alice.New(loggingHandler, handlers.CompressHandler)

	// http.Handle("/", middlewareFirst(middlewareSecond(indexHandler)))
	// http.Handle("/message", middlewareFirst(middlewareSecond(messageHandler)))

	http.Handle("/", commonHandlers.ThenFunc(indexHandler))
	http.Handle("/message", commonHandlers.ThenFunc(messageHandler))

	server := &http.Server{
		Addr: ":8001",
	}

	log.Println("Listening...")
	server.ListenAndServe()
}
