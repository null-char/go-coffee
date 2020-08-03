// Package docs consists of business logic to serve and set up our swagger docs.
package docs

import (
	"github.com/gin-gonic/gin"
	openapiMiddleware "github.com/go-openapi/runtime/middleware"
)

// RegisterRoutes registers all routes related to swagger docs. Calling this function will
// also register a basic file server for serving our swagger-docs.yaml file. The docs will be
// served using ReDoc in the relative path /docs.
func RegisterRoutes(r *gin.RouterGroup) {
	r.StaticFile("swagger.yaml", "./swagger-docs.yaml")
	redocOpts := openapiMiddleware.RedocOpts{SpecURL: "/swagger.yaml"}
	rh := openapiMiddleware.Redoc(redocOpts, nil)
	r.GET("/docs", gin.WrapH(rh))
}
