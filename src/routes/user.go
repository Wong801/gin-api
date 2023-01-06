package route

import (
	controller "github.com/Wong801/gin-api/src/controllers"
	"github.com/gin-gonic/gin"
)

func (r handler) addUsers(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")
	userController := controller.InitUserController()

	userRoute.POST("/register", userController.Register())
	userRoute.POST("/login", userController.Login())
}
