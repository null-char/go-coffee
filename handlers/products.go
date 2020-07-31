package handlers

import (
	"log"
	"net/http"

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
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (p *Products) getProducts(rw http.ResponseWriter) {
	productsList := data.GetProducts()
	productsList.ToJSON(rw, productsList)
}
