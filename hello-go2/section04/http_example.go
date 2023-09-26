package section04

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

type ShortenRes struct {
	AppId         string    `json:"app_id,omitempty"`
	AppName       string    `json:"app_name,omitempty"`
	SourceUrl     string    `json:"source_url,omitempty"`
	ShortUrl      string    `json:"short_url,omitempty"`
	Expiration    time.Time `json:"expiration,omitempty"`
	ExpirationStr string    `json:"expiration_str,omitempty"`
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

	now := time.Now()
	payload := ShortenRes{
		AppName:       "测试应用",
		AppId:         "appId123",
		SourceUrl:     "https://www.baidu.com",
		ShortUrl:      "t.cn/abc",
		Expiration:    now,
		ExpirationStr: now.Format("2006-01-02 15:04:05"),
		// 时间格式化
	}

	respondWithJSON(w, http.StatusOK, payload)

	// fmt.Fprintf(w, "welcome")

}

func message(w http.ResponseWriter, r *http.Request) {
	log.Println("Executed message Handler")
	_, _ = fmt.Fprintf(w, "message...")
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
	_ = server.ListenAndServe()
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	resp, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(resp)
}
