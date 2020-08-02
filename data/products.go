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
func AddProduct(newProduct *Product) {
	newProduct.ID = nextID()
	newProduct.CreatedOn = time.Now().UTC().String()
	productsList = append(productsList, newProduct)
}

// Totally how DBs work
func nextID() int {
	lp := productsList[len(productsList)-1]
	return lp.ID + 1
}

// UpdateProduct receives an id and the newData struct and replaces the required product with
// the newly provided data. It will either return a pointer to a Product struct or a null pointer
// if no product could be found.
func UpdateProduct(id int, newData Product) *Product {
	prodToUpdate, _, err := FindProduct(id)
	if err != nil {
		return nil
	}

	// Grab the value of CreatedOn before we replace
	createdOn := prodToUpdate.CreatedOn

	*prodToUpdate = newData
	prodToUpdate.ID = id
	prodToUpdate.CreatedOn = createdOn
	prodToUpdate.UpdatedOn = time.Now().UTC().String()

	return prodToUpdate
}

// DeleteProduct takes an id and deletes the product with that id if found. If the product
// does not exist, it'll return an error.
func DeleteProduct(id int) error {
	_, i, err := FindProduct(id)
	if err != nil {
		return err
	}

	productsList = removeIndex(productsList, i)
	return nil
}

func removeIndex(slice ProductsList, i int) ProductsList {
	return append(slice[0:i], slice[i+1:]...)
}

// FindProduct finds and returns the product with the given id including its index in the slice.
func FindProduct(id int) (*Product, int, error) {
	for i, prod := range productsList {
		if prod.ID == id {
			return prod, i, nil
		}
	}

	// Error out if we couldn't find a product with the specified id
	err := fmt.Errorf("product with id %v not found", id)
	return nil, -1, err
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
