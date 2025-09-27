package middleware

import (
	"net/http"
	"strings"
	"web3-go-blog/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	CtxUserID   = "user_id"
	CtxUsername = "username"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": -1, "message": "Unauthorized"})
			return
		}
		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": -1, "message": "invalid auth header"})
			return
		}
		token, err := utils.ParseToken(parts[1])
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": -1, "message": "invalid token"})
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		c.Set(CtxUserID, uint(claims["sub"].(float64)))
		c.Set(CtxUsername, claims["usr"].(string))
		c.Next()
	}
}
