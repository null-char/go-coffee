package products

import (
	"github.com/gin-gonic/gin"
	"github.com/null-char/go-coffee/middleware"
)

// RegisterRoutes receives a RouterGroup and attaches all the appropriate request handlers
// pertaining to products.
func RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/", getProducts)
	r.POST("/", addProduct)

	sr := r.Group("/:id").Use(middleware.ParseInt("id"))
	sr.PUT("/", updateProduct)
	sr.DELETE("/", deleteProduct)
}
