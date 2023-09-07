package product_handlers

import (
	"github.com/p97k/on-mark/data"
	"net/http"
)

func (p *Products) GetProducts(response http.ResponseWriter, request *http.Request) {
	p.l.Println("handle GET product")

	lp := data.GetProducts()
	err := lp.ToJSON(response)
	if err != nil {
		http.Error(response, "unable to convert json!", http.StatusInternalServerError)
	}
}
