package middleware

import (
	"cat-slave/pkg/http/result"
	"cat-slave/pkg/token"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		if _, err := token.ParseRequest(c); err != nil {
			result.Error(c, err)
			c.Abort()
			return
		}
		c.Next()
	}
}
