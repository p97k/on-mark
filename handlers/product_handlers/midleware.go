// Package classification of Product API
//
// Documentation for Product API
//
// Schemes http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
//
//Products:
// -application/json
// swagger:meta

package product_handlers

import (
	"context"
	"fmt"
	"github.com/p97k/on-mark/datas"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

type KeyProduct struct{}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		prod := datas.Product{}
		err := prod.FromJSON(request.Body)
		if err != nil {
			http.Error(response, "Ops, unable to create json", http.StatusBadRequest)
			return
		}

		err = prod.Validate()
		if err != nil {
			http.Error(response, fmt.Sprintf("Error Validating Product: %s", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(request.Context(), KeyProduct{}, prod)
		request = request.WithContext(ctx)

		next.ServeHTTP(response, request)
	})
}
