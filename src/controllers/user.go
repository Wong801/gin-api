package controller

import (
	"net/http"

	entity "github.com/Wong801/gin-api/src/entities"
	service "github.com/Wong801/gin-api/src/services"
	"github.com/gin-gonic/gin"
)

func UserRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body entity.User
		if err := c.ShouldBindJSON(&body); err != nil {
			c.Set("status", http.StatusBadRequest)
			c.Set("error", service.MakeRequestError(err))
			return
		}
		c.Set("status", http.StatusOK)
		c.Set("data", body)
		c.Next()
	}
}
