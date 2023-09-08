package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	goHandler "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	dataAccess "github.com/p97k/on-mark/data-access"
	"github.com/p97k/on-mark/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	tempLog := log.New(os.Stdout, "on-mark", log.LstdFlags)

	//connecting to database
	dataAccess.Start()

	serveMux := mux.NewRouter()

	routes.InitRoutes(serveMux)

	origins := []string{"http://localhost:3000"}
	corsHandler := goHandler.CORS(goHandler.AllowedOrigins(origins))

	//create a new server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      corsHandler(serveMux),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	//start the server
	go func() {
		tempLog.Println("On Mark is up and running :)")
		startSeverErr := server.ListenAndServe()
		if startSeverErr != nil {
			tempLog.Fatal(startSeverErr)
		}
	}()

	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	sig := <-signalChannel
	tempLog.Println("gracefully shutdown!", sig)

	timeContext, _ := context.WithTimeout(context.Background(), 30*time.Second)

	shutDownErr := server.Shutdown(timeContext)
	if shutDownErr != nil {
		return
	}
}
