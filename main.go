package main

import (
	"github.com/p97k/on-mark/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	tempLog := log.New(os.Stdout, "product-api", log.LstdFlags)
	handler := handlers.NewLogger(tempLog)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", handler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
