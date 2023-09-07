package routes

import "github.com/gorilla/mux"

func InitRoutes(serveMux *mux.Router) {
	ProductRoutes(serveMux)
}
