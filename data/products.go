package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Product defines the structure of a product in our API.
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:""`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// ProductsList is a type that is a slice of pointers to Product struct.
type ProductsList []*Product

// ToJSON writes a JSON encoded string into the provided writer.
func (p *ProductsList) ToJSON(writer io.Writer, v interface{}) {
	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(v); err != nil {
		writer.Write([]byte("Status Code: 500 \n Internal server error. Please try again."))
		return
	}
}

// GetProducts returns a slice of Products
func GetProducts() ProductsList {
	return productsList
}

var productsList = ProductsList{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       3.0,
		SKU:         "hg43i",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Latte",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "ab3cr",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
