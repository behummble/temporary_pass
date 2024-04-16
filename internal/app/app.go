package app

import (
	"net/http"
	"time"
	"log/slog"
	
)

var (
	log *slog.Logger
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

}

func startTokenValidationProccess() {

}

func initLog() {
	log = slog.New()
}

func fillHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", )
}