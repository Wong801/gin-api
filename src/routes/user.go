package route

import (
	controller "github.com/Wong801/gin-api/src/controllers"
	"github.com/gin-gonic/gin"
)

func (r handler) addUsers(rg *gin.RouterGroup) {
	users := rg.Group("/user")

	users.POST("/register", controller.UserRegister())
	users.POST("/login", controller.UserLogin())
}
