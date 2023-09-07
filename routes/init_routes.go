package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

func InitRoutes(serveMux *mux.Router) {
	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		_, err := response.Write([]byte("This is the home of On-Mark!, Welcome :)"))
		if err != nil {
			return
		}
	})

	ProductRoutes(serveMux)
}
