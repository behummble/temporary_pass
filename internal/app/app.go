package app

import (
	"net/http"
	"time"
	"log"
	"github.com/behummble/temporary_pass/internal/handlers"
	"github.com/behumble/temprorary_pass/internal/service"
)


func Run() {
	initLog()
	startHandlers()
	go startTokenValidationProccess()
}

func startHandlers() {

	mux := http.NewServeMux()
	
	fillHandlers(mux)
	
	server := &http.Server{
		Addr: "localhost:8080",
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

func startTokenValidationProccess() {
	
}

func initLog() {
//	log = log.New(os.Stdout, "TEMP_PASS", 1)
}

func fillHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", handlers.RedisHandler)
}