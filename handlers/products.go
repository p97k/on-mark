package handlers

import (
	"github.com/p97k/on-mark/products"
	"log"
	"net/http"
	"regexp"
	"strconv"
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

	if request.Method == http.MethodPost {
		p.addProducts(response, request)
		return
	}

	if request.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)
		finalUrl := reg.FindAllStringSubmatch(request.URL.Path, -1)

		if len(finalUrl) != 1 {
			p.l.Println("invalid URI more than one id")
			http.Error(response, "Invalid URL", http.StatusBadRequest)
			return
		}

		if len(finalUrl[0]) != 2 {
			p.l.Println("invalid URI more than capture group id")
			http.Error(response, "Invalid URL", http.StatusBadRequest)
			return
		}

		idString := finalUrl[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("invalid URI unable to convert to num", idString)
			http.Error(response, "Invalid URL", http.StatusBadRequest)
			return
		}

		p.updateProduct(id, response, request)
		return
	}

	response.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(response http.ResponseWriter, request *http.Request) {
	p.l.Println("handle GET product")

	lp := products.GetProducts()
	err := lp.ToJSON(response)
	if err != nil {
		http.Error(response, "unable to convert json!", http.StatusInternalServerError)
	}
}

func (p *Products) addProducts(response http.ResponseWriter, request *http.Request) {
	p.l.Println("handle POST product")

	prod := &products.Product{}
	err := prod.FromJSON(request.Body)
	if err != nil {
		http.Error(response, "Ops, unable to create json", http.StatusBadRequest)
	}

	products.AddProduct(prod)

	p.l.Printf("Product: %#v", prod)
}

func (p *Products) updateProduct(id int, response http.ResponseWriter, request *http.Request) {
	p.l.Println("handle PUT product")

	prod := &products.Product{}
	err := prod.FromJSON(request.Body)
	if err != nil {
		http.Error(response, "Ops, unable to create json", http.StatusBadRequest)
	}

	err = products.UpdateProduct(id, prod)
	if err == products.ErrProductNotFound {
		http.Error(response, "product not found", http.StatusNotFound)
	}

	if err != nil {
		http.Error(response, "something wrong", http.StatusInternalServerError)
	}
}
