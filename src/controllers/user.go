package controller

import (
	"net/http"

	"github.com/Wong801/gin-api/src/api"
	model "github.com/Wong801/gin-api/src/models"
	"github.com/gin-gonic/gin"
)

func (uc UserController) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body model.User
		if err := c.ShouldBindJSON(&body); err != nil {
			c.Set("status", http.StatusBadRequest)
			c.Set("error", api.MakeRequestError(err))
			return
		}
		status, err := uc.s.Register(&body)
		c.Set("status", status)
		if err != nil {
			c.Set("error", api.MakeResultError(err))
			return
		}
		c.Set("data", map[string]string{
			"message": "Register Success",
		})
		c.Next()
	}
}

func (uc UserController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body model.UserLogin
		if err := c.ShouldBindJSON(&body); err != nil {
			c.Set("status", http.StatusBadRequest)
			c.Set("error", api.MakeRequestError(err))
			return
		}
		status, token, err := uc.s.Login(&body)
		c.Set("status", status)
		if err != nil {
			c.Set("error", api.MakeResultError(err))
			return
		}
		c.SetCookie("jwt", token.Jwt, token.MaxAge, "/", token.Domain, token.Secure, token.HttpOnly)
		c.Set("data", map[string]string{
			"message": "Login Success",
		})
		c.Next()
	}
}

func (uc UserController) ChangePassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body model.UserChangePassword
		if err := c.ShouldBindJSON(&body); err != nil {
			c.Set("status", http.StatusBadRequest)
			c.Set("error", api.MakeRequestError(err))
			return
		}
		status, err := uc.s.ChangePassword(c.MustGet("user_id").(int), &body)
		c.Set("status", status)
		if err != nil {
			c.Set("error", api.MakeResultError(err))
			return
		}
		c.Set("data", map[string]string{
			"message": "Change Password Success",
		})
		c.Next()
	}
}
