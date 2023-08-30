package main

import (
	"context"
	"github.com/p97k/on-mark/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	tempLog := log.New(os.Stdout, "product-api", log.LstdFlags)
	//handler := handlers.NewLogger(tempLog)
	productHandler := handlers.NewProducts(tempLog)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", productHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			tempLog.Fatal(err)
		}
	}()

	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	sig := <-signalChannel
	tempLog.Println("gracefully shutdown!", sig)

	timeContext, _ := context.WithTimeout(context.Background(), 30*time.Second)

	err := server.Shutdown(timeContext)
	if err != nil {
		return
	}
}
