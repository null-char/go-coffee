// Package classification Products API
//
// The products APIs allows consumers to get a list of products (mostly just coffee), add to
// them and update / delete a certain product.
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//	License: MIT http://opensource.org/licenses/MIT
//	Contact: null-char<skp.frl@gmail.com> https://nullchar.now.sh
//
//	Consumes:
//     - application/json
//
//  Produces:
//     - application/json
//
// swagger:meta
package products

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/null-char/go-coffee/data"
)

// swagger:route GET /products products listProducts
//
// Lists all products currently in the data store.
//
// You can currently only get all of the products by default.
// Getting by id is not supported yet.
//
//	Produces:
//	- application/json
//
//	Schemes: http
//
//	Responses:
//	  200: productsListResponse
func getProducts(ctx *gin.Context) {
	productsList := data.GetProducts()
	ctx.JSON(http.StatusOK, productsList)
}

// swagger:route POST /products products addProduct
//
// Adds a product to the data store.
//
// Once the product is added, you'll get a response of 201 with the newly created product.
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Schemes: http
//
//	Responses:
//	  201: productsListResponse
//
// addProduct adds a product into the DB
func addProduct(ctx *gin.Context) {
	var prod data.Product
	if err := ctx.ShouldBindJSON(&prod); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// If binding is successful, then we add the product to our fake DB
	data.AddProduct(&prod)
	ctx.JSON(http.StatusCreated, prod)
}

// swagger:route PUT /products/{id} products updateProduct
//
// Updates a product with the specified ID.
//
// The server will respond with a status of 200 if the product was successfully updated.
// If there exists no such product with the given id, the server will respond with a 404.
//
//	Produces:
//	- application/json
//
//	Consumes:
//	- application/json
//
//	Schemes: http
//
//	Responses:
//	  200: productResponse
//	  404:
//	  400: errorResponse
func updateProduct(ctx *gin.Context) {
	var newData data.Product
	if err := ctx.ShouldBindJSON(&newData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assert that id is an int. Should panic if not.
	id := ctx.MustGet("id").(int)
	// updatedProd will be nil if we couldn't find the required product
	if updatedProd := data.UpdateProduct(id, newData); updatedProd != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Updated product with id %v", id), "product": updatedProd,
		})

		return
	}
	ctx.Status(http.StatusNotFound)
}

// swagger:route DELETE /products/{id} products deleteProduct
//
// Deletes a product with the specified ID from the data store.
//
// The server will respond with a status of 204 if the delete was sucessful.
// If there exists no such product with the given id, the server will respond with a 404.
//
//	Schemes: http
//
//	Responses:
//	  204:
//	  404:
//
// deleteProduct deletes a product from the DB.
func deleteProduct(ctx *gin.Context) {
	id := ctx.MustGet("id").(int)
	data.DeleteProduct(id)
	ctx.Status(http.StatusNoContent)
}
