package middleware

import (
	"net/http"

	entity "github.com/Wong801/gin-api/src/entities"
	"github.com/gin-gonic/gin"
)

func Recovery() func(c *gin.Context) {
	return gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &entity.HttpResponse{
			Success: false,
			Data: map[string]string{
				"message": err.(string),
			},
		})
	})
}
