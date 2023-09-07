package product_handlers

import (
	"github.com/p97k/on-mark/data"
	"net/http"
)

func (p *Products) AddProducts(response http.ResponseWriter, request *http.Request) {
	p.l.Println("handle POST product")

	prod := request.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&prod)

	p.l.Printf("Product: %#v", prod)
}
