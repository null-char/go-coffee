// responses describes the shape of our responses.

package products

import "github.com/null-char/go-coffee/data"

// productsListResponse contains a body with a list of products from the data store for GET /products
// swagger:response
type productsListResponse struct {
	// A list of Products from the data store
	// in: body
	Body []data.Product
}

// productResponse describes the response for operations such as POST, GET and PUT for /products
// swagger:response
type productResponse struct {
	// A product from the data store with id, name, description, price and sku.
	// in: body
	Body data.Product
}

// errorResponse is what the server will respond with in case it's met with validation
// errors or any other generic error.
// swagger:response
type errorResponse struct {
	// The response body will contain a single "error" field which is a string that describes
	// the error be it a validation error or anything else.
	// in: body
	Body struct {
		// required: true
		Error string `json:"error"`
	}
}
