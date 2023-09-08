package datas

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	data_access "github.com/p97k/on-mark/data-access"
	"io"
	"log"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	CreatedAt   string  `json:"-"`
	UpdatedAt   string  `json:"-"`
}

type Products []*Product

func (p *Product) FromJSON(reader io.Reader) error {
	e := json.NewDecoder(reader)
	return e.Decode(p)
}

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = generateNextId()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}

func DeleteProduct(id int) error {
	_, i, _ := findProduct(id)

	if i == -1 {
		return ErrProductNotFound
	}

	productList = append(productList[:i], productList[i+1])

	return nil
}

var ErrProductNotFound = fmt.Errorf("product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func generateNextId() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Milk + Espresso",
		Price:       2.99,
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Strong coffee without milk",
		Price:       1.99,
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
}

//Migrating To Mysql

func GetProductList() error {
	var product Product

	sqlQuery := `SELECT * FROM products`

	rows, err := data_access.DB.Query(sqlQuery)
	if err != nil {
		return err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		err := rows.Scan(product)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(product)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
