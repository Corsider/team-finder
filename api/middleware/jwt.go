package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"team-finder/domain"
	"team-finder/internal/utils"
)

func JWT(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		t := strings.Split(auth, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := utils.IsAuthorized(authToken, secret)
			if authorized {
				userId, err1 := utils.ExtractID(authToken, secret)
				if err1 != nil {
					c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err1.Error()})
					c.Abort()
					return
				}
				c.Set("x-user-id", userId)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Unauthorized"})
		c.Abort()
	}
}
