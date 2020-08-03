// params.go containes the swagger documentation for goswagger codegen.

package products

// swagger:parameters updateProduct deleteProduct
type productIDParam struct {
	// The ID of the product.
	// in: path
	// required: true
	// unique: true
	// min: 0
	ID int `json:"id"`
}
