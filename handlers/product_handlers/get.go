package product_handlers

import (
	"github.com/p97k/on-mark/datas"
	"net/http"
)

func (p *Products) GetProducts(response http.ResponseWriter, request *http.Request) {
	p.l.Println("handle GET product")

	//err := datas.GetProductList()
	//if err != nil {
	//	return
	//}

	lp := datas.GetProducts()

	err := lp.ToJSON(response)
	if err != nil {
		http.Error(response, "unable to convert json!", http.StatusInternalServerError)
	}
}
