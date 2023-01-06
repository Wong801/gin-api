package controller

import service "github.com/Wong801/gin-api/src/services"

type UserController struct {
	s service.UserService
}

func InitUserController() *UserController {
	us := &UserController{
		s: *service.InitUserService(),
	}
	return us
}
