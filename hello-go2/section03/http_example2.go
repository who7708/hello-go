package section03

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"log"
	"net/http"
	"os"
)

func loggingHandler(next http.Handler) http.Handler {

	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}

	return handlers.LoggingHandler(logFile, next)
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("executing index handler")
	// w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "welcome")
	// w.WriteHeader(code)
	// w.Write(resp)
}

func message(w http.ResponseWriter, r *http.Request) {
	log.Println("Executed message Handler")
	fmt.Fprintf(w, "message...")
}

func iconHandler(w http.ResponseWriter, r *http.Request) {

}

func Run() {
	router := mux.NewRouter()
	m := alice.New(loggingHandler, handlers.CompressHandler)

	router.Handle("/favicon.ico", m.ThenFunc(iconHandler)).Methods("POST")
	router.Handle("/", m.ThenFunc(index)).Methods("GET")
	router.Handle("/index", m.ThenFunc(index)).Methods("GET")
	router.Handle("/message", m.ThenFunc(message)).Methods("GET")

	server := &http.Server{
		Addr:    ":8001",
		Handler: router,
	}

	log.Println("Listening...")
	server.ListenAndServe()
}
