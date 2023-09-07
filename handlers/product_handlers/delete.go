package product_handlers

import (
	"github.com/gorilla/mux"
	"github.com/p97k/on-mark/data"
	"net/http"
	"strconv"
)

func (p *Products) DeleteProduct(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(response, "unable to convert id!", http.StatusBadRequest)
		return
	}

	p.l.Println("handle DELETE product", id)

	err = data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(response, "Product Not Found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(response, "Product Not Found ( Server Issue )", http.StatusInternalServerError)
		return
	}
}
