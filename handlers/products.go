package handlers

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/p97k/on-mark/products"
	"log"
	"net/http"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(response http.ResponseWriter, request *http.Request) {
	p.l.Println("handle GET product")

	lp := products.GetProducts()
	err := lp.ToJSON(response)
	if err != nil {
		http.Error(response, "unable to convert json!", http.StatusInternalServerError)
	}
}

func (p *Products) AddProducts(response http.ResponseWriter, request *http.Request) {
	p.l.Println("handle POST product")

	prod := request.Context().Value(KeyProduct{}).(products.Product)

	products.AddProduct(&prod)

	p.l.Printf("Product: %#v", prod)
}

func (p *Products) UpdateProduct(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	println(vars)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(response, "unable to convert id!", http.StatusBadRequest)
		return
	}

	p.l.Println("handle PUT product", id)

	prod := request.Context().Value(KeyProduct{}).(products.Product)

	err = products.UpdateProduct(id, &prod)
	if err == products.ErrProductNotFound {
		http.Error(response, "product not found", http.StatusNotFound)
	}

	if err != nil {
		http.Error(response, "something wrong", http.StatusInternalServerError)
	}
}

type KeyProduct struct{}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		prod := products.Product{}
		err := prod.FromJSON(request.Body)
		if err != nil {
			http.Error(response, "Ops, unable to create json", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(request.Context(), KeyProduct{}, prod)
		request = request.WithContext(ctx)

		next.ServeHTTP(response, request)
	})
}
