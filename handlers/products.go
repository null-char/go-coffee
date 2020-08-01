package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/null-char/go-coffee/data"
)

// Products handler struct
type Products struct {
	l *log.Logger
}

// NewProducts returns a Products handler for our REST API
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.getProducts(rw)
	case http.MethodPost:
		p.postProducts(r.Body, rw)
	case http.MethodPut:
		p.updateProduct(rw, r.Body, r.URL.Path)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (p *Products) getProducts(rw http.ResponseWriter) {
	p.l.Println("GET Products")
	productsList := data.GetProducts()
	productsList.ToJSON(rw, productsList)
}

func (p *Products) postProducts(body io.ReadCloser, rw http.ResponseWriter) {
	p.l.Println("POST Product")
	newProduct := new(data.Product)
	if err := newProduct.FromJSON(body); err != nil {
		p.l.Println("Invalid request body")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(`{"message": "Invalid request body"}`))
		return
	}

	data.AddProduct(*newProduct)
	rw.WriteHeader(http.StatusCreated)
	p.l.Printf("Created new product: %+v", *newProduct)
	rw.Write([]byte(`{"message": "Successfully created new product"}`))
}

func (p *Products) updateProduct(rw http.ResponseWriter, body io.ReadCloser, URI string) {
	p.l.Println("PUT Product")

	// First, we want to pull out the ID from the URI
	rgxp := regexp.MustCompile(`/{[0-9]+}`)
	id, err := strconv.ParseInt(rgxp.FindStringSubmatch(URI)[0][1:], 10, 0)

	// Return a bad request status code if id could not be parsed to int
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(`{"message": "Invalid id"}`))
		return
	}

	// Decode body
	newProd := new(data.Product)
	newProd.FromJSON(body)

	// Find and update the required product
	data.UpdateProduct(int(id), *newProd)
	rw.Write([]byte(fmt.Sprintf(`{"message": "Successfully updated the product with id %v"}`, id)))
}
