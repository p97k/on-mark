package product_handlers

import (
	"github.com/gorilla/mux"
	"github.com/p97k/on-mark/data"
	"net/http"
	"strconv"
)

func (p *Products) UpdateProduct(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(response, "unable to convert id!", http.StatusBadRequest)
		return
	}

	p.l.Println("handle PUT product", id)

	prod := request.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(response, "product not found", http.StatusNotFound)
	}

	if err != nil {
		http.Error(response, "something wrong", http.StatusInternalServerError)
	}
}
