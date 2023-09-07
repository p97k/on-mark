package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/p97k/on-mark/handlers/product_handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	routeIdRegex := "/{id:[0-9]+}"

	tempLog := log.New(os.Stdout, "product-api", log.LstdFlags)
	productHandler := product_handlers.NewProducts(tempLog)

	serveMux := mux.NewRouter()

	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", productHandler.GetProducts)

	putRouter := serveMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc(routeIdRegex, productHandler.UpdateProduct)
	putRouter.Use(productHandler.MiddlewareProductValidation)

	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", productHandler.AddProducts)
	postRouter.Use(productHandler.MiddlewareProductValidation)

	deleteRouter := serveMux.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc(routeIdRegex, productHandler.DeleteProduct)

	//create a new server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      serveMux,
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

	err := server.Shutdown(timeContext)
	if err != nil {
		return
	}
}
