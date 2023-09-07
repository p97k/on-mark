package routes

import (
	"github.com/gorilla/mux"
	"github.com/p97k/on-mark/handlers/product_handlers"
	"log"
	"net/http"
	"os"
)

func ProductRoutes(serveMux *mux.Router) {
	baseProductUrl := "/v1/products/"
	routeIdRegex := "{id:[0-9]+}"
	idRequiredRoutes := baseProductUrl + routeIdRegex

	tempLog := log.New(os.Stdout, "product-api", log.LstdFlags)

	productHandler := product_handlers.NewProducts(tempLog)

	//GET
	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc(baseProductUrl, productHandler.GetProducts)

	//PUT
	putRouter := serveMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc(idRequiredRoutes, productHandler.UpdateProduct)
	putRouter.Use(productHandler.MiddlewareProductValidation)

	//POST
	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc(baseProductUrl, productHandler.AddProducts)
	postRouter.Use(productHandler.MiddlewareProductValidation)

	//DELETE
	deleteRouter := serveMux.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc(idRequiredRoutes, productHandler.DeleteProduct)
}
