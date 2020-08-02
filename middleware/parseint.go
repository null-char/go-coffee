package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ParseInt is a middleware that takes in a paramKey, looks it up in the provided context, then
// parses it to an int. It finally sets it as a key value pair (where the key is paramKey)
// in its provided context.
func ParseInt(paramKey string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param := ctx.Param(paramKey)
		i, err := strconv.Atoi(param)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("%s should be a number", paramKey),
			})
			return
		}

		ctx.Set(paramKey, i)
		ctx.Next()
	}
}
