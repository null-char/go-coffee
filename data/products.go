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
	Name        string  `json:"name" binding:"required,min=5"`
	Description string  `json:"description" binding:"required,min=15" `
	Price       float32 `json:"price" binding:"required,gt=0"`
	SKU         string  `json:"sku" binding:"required,len=5"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// FromJSON initializes a product struct decoded from JSON. Returns an error if not successful.
func (p *Product) FromJSON(body io.ReadCloser) error {
	decoder := json.NewDecoder(body)
	return decoder.Decode(p)
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

// AddProduct adds a new product to the products slice
func AddProduct(newProduct Product) {
	productsList = append(productsList, &newProduct)
}

// UpdateProduct receives an id and the newData struct and replaces the required product with
// the newly provided data.
func UpdateProduct(id int, newData Product) (*Product, error) {
	prodToUpdate, err := FindProduct(id)
	if err != nil {
		return nil, err
	}

	*prodToUpdate = newData
	return prodToUpdate, nil
}

// FindProduct finds and returns the product with the given id.
func FindProduct(id int) (*Product, error) {
	for _, prod := range productsList {
		if prod.ID == id {
			return prod, nil
		}
	}

	// Error out if we couldn't find a product with the specified id
	err := fmt.Errorf("product with id %v not found", id)
	return nil, err
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
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "ab3cr",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
