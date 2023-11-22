package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"team-finder/domain"
	"team-finder/internal/utils"
)

func JWT(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			authorized, err := utils.IsAuthorized(c.GetHeader("Authorization"), secret)
			if authorized {
				c.Next()
			}
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Error: err.Error()})
			c.Abort()
		} else {
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "Unauthorized"})
			c.Abort()
		}
	}
}
