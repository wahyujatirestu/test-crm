package middleware

import (
	"strings"
	"net/http"

	"github.com/gin-gonic/gin"
	"test_crm/utils"
)

type AuthMiddleware interface {
	RequireToken() gin.HandlerFunc
}

type authMiddleware struct {
	jwt utils.JWTService
}

func NewAuthMiddleware(jwt utils.JWTService) AuthMiddleware {
	return &authMiddleware{jwt}
}

func (a *authMiddleware) RequireToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(header, "Bearer ")
		if _, err := a.jwt.ValidateToken(token); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
