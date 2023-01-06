package middleware

import (
	"net/http"
	"strconv"

	"github.com/Wong801/gin-api/src/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func (m Middleware) Authenticate() func(c *gin.Context) {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("jwt")

		if err != nil {
			c.Set("status", http.StatusUnauthorized)
			c.Set("error", "Unauthorized")
			return
		}

		claims := &jwt.RegisteredClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetEnv("JWT_SECRET", "secret")), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.Set("status", http.StatusUnauthorized)
				c.Set("error", "Unauthorized")
				return
			}
			c.Set("status", http.StatusBadRequest)
			c.Set("error", "Unauthorized")
			return
		}

		if !token.Valid {
			c.Set("status", http.StatusUnauthorized)
			c.Set("error", "Invalid Token")
			return
		}
		id, _ := strconv.Atoi(claims.ID)

		c.Set("user_id", id)
		c.Next()
	}
}
