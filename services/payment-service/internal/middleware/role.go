/**
 * Role-based authorization middleware.
 */
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetString("role") != role {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}
