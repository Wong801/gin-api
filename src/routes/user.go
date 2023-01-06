package route

import (
	controller "github.com/Wong801/gin-api/src/controllers"
	middleware "github.com/Wong801/gin-api/src/middlewares"
	"github.com/gin-gonic/gin"
)

func (r handler) addUsers(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")
	userController := controller.InitUserController()
	m := middleware.InitMiddleware()

	userRoute.POST("/register", userController.Register())
	userRoute.POST("/login", userController.Login())

	userRoute.Use(m.Authenticate())

	userRoute.PATCH("/change-password", userController.ChangePassword())
}
