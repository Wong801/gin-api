package route

import (
	controller "github.com/Wong801/gin-api/src/controllers"
	middleware "github.com/Wong801/gin-api/src/middlewares"
	"github.com/gin-gonic/gin"
)

func (r route) addUsers(rg *gin.RouterGroup) {
	users := rg.Group("/user")

	users.POST("/register", controller.UserRegister(), middleware.Response())
}
