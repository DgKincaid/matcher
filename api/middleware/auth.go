package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Auth middleware currently empty but would be used to verify the jwt making sure the are authenticated.
// Authorization would be taken care of at the endpoints due to the need for context.
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Auth middleware")
	}
}
