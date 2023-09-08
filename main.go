package main

import (
	"context"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	goHandler "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/p97k/on-mark/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	tempLog := log.New(os.Stdout, "on-mark", log.LstdFlags)

	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "on-mark-db",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	tempLog.Println("Successfully Connected to on-mark DB!")

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

	err = server.Shutdown(timeContext)
	if err != nil {
		return
	}
}
