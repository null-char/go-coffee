package products

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/null-char/go-coffee/data"
)

func getProducts(ctx *gin.Context) {
	productsList := data.GetProducts()
	ctx.JSON(http.StatusOK, productsList)
}

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
		ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Updated product with id %v", id), "product": updatedProd})
		return
	}
	ctx.Status(http.StatusNotFound)
}

func deleteProduct(ctx *gin.Context) {
	id := ctx.MustGet("id").(int)
	data.DeleteProduct(id)
	ctx.Status(http.StatusOK)
}
