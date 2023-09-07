package product_handlers

import (
	"github.com/p97k/on-mark/datas"
	"net/http"
)

func (p *Products) AddProducts(response http.ResponseWriter, request *http.Request) {
	p.l.Println("handle POST product")

	prod := request.Context().Value(KeyProduct{}).(datas.Product)

	datas.AddProduct(&prod)

	p.l.Printf("Product: %#v", prod)
}
