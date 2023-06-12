package controller

import (
	"net/http"

	"github.com/Wong801/gin-api/src/api"
	model "github.com/Wong801/gin-api/src/models"
	service "github.com/Wong801/gin-api/src/services"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

type UserController struct {
	s service.UserService
}

func InitUserController() *UserController {
	uc := &UserController{
		s: *service.InitUserService(),
	}
	return uc
}

func (uc UserController) GetProfile() gin.HandlerFunc {
	return func(c *gin.Context) {

		status, data, err := uc.s.GetUser()
		uc.s.DB.Close()
		c.Set("status", status)
		if err != nil {
			c.Set("error", api.MakeResultError((err)))
			return
		}
		c.Set("data", data)
		c.Next()
	}
}

func (uc UserController) UpdateProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body model.UserBase
		if err := c.ShouldBind(&body); err != nil {
			c.Set("status", http.StatusBadRequest)
			c.Set("error", api.MakeRequestError(err))
		}
		status, data, err := uc.s.UpdateUser(c.MustGet("user_id").(int), &body)
		uc.s.DB.Close()

		c.Set("status", status)
		if err != nil {
			c.Set("error", api.MakeResultError(err))
			return
		}
		c.Set("data", data)
		c.Next()
	}
}

func (uc UserController) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body model.User
		if err := c.ShouldBind(&body); err != nil {
			c.Set("status", http.StatusBadRequest)
			c.Set("error", api.MakeRequestError(err))
			return
		}
		status, err := uc.s.Register(&body)
		uc.s.DB.Close()
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
		if err := c.ShouldBind(&body); err != nil {
			c.Set("status", http.StatusBadRequest)
			c.Set("error", api.MakeRequestError(err))
			return
		}
		status, token, err := uc.s.Login(&body)
		uc.s.DB.Close()
		c.Set("status", status)
		if err != nil {
			c.Set("error", api.MakeResultError(err))
			return
		}
		c.SetCookie("jwt", token.Jwt, token.MaxAge, "/", token.Domain, token.Secure, token.HttpOnly)
		c.SetCookie("X-CSRF-TOKEN", csrf.GetToken(c), token.MaxAge, "/", token.Domain, false, false)
		c.Set("data", map[string]string{
			"message": "Login Success",
		})
		c.Next()
	}
}

func (uc UserController) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("jwt", "", -1, "/", "", true, true)
		c.SetCookie("X-CSRF-TOKEN", "", -1, "/", "", false, false)
		c.Next()
	}
}

func (uc UserController) ChangePassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body model.UserChangePassword
		if err := c.ShouldBind(&body); err != nil {
			c.Set("status", http.StatusBadRequest)
			c.Set("error", api.MakeRequestError(err))
			return
		}
		status, err := uc.s.ChangePassword(c.MustGet("user_id").(int), &body)
		uc.s.DB.Close()
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
