package middleware

import (
	"net/http"

	entity "github.com/Wong801/gin-api/src/entities"
	"github.com/gin-gonic/gin"
)

func getErrorStatus(status any) int {
	if status != nil {
		return status.(int)
	}
	return http.StatusInternalServerError
}

func getSuccessStatus(status any) int {
	if status != nil {
		return status.(int)
	}
	return http.StatusOK
}

func getErrorMessage(c *gin.Context) any {
	err := c.Errors.Last()
	if err != nil {
		return entity.ResultError{Reason: c.Error(err).Error()}
	}
	customError, _ := c.Get("error")
	return customError
}

func (m middleware) Response() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := getErrorMessage(c)
		status, _ := c.Get("status")
		if err != nil {
			c.AbortWithStatusJSON(getErrorStatus(status), &entity.HttpResponse{
				Success: false,
				Data:    err,
			})
			return
		}
		c.JSON(getSuccessStatus(status), &entity.HttpResponse{
			Success: true,
			Data:    c.MustGet("data"),
		})
	}
}
