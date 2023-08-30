package handlers

import (
	"github.com/p97k/on-mark/products"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		p.getProducts(response, request)
		return
	}

	response.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(response http.ResponseWriter, request *http.Request) {
	lp := products.GetProducts()
	err := lp.ToJSON(response)
	if err != nil {
		http.Error(response, "unable to convert json!", http.StatusInternalServerError)
	}
}
