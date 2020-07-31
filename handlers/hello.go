package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// The Hello handler struct
type Hello struct {
	l *log.Logger
}

// NewHello creates a new Hello struct.
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// ServeHTTP handles serve logic for "/" path
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	logger := h.l
	logger.Println("Hello World")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Bad Request", http.StatusBadRequest)
		return
	}

	logger.Printf("Data is: %s", d)
	fmt.Fprintf(rw, "Hello %s", d)
}
