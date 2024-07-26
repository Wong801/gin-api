package route

import (
	controller "github.com/Wong801/gin-api/src/controllers"
	middleware "github.com/Wong801/gin-api/src/middlewares"
	"github.com/gin-gonic/gin"
)

func (h handler) addSkills(rg *gin.RouterGroup) {
	skillRoute := rg.Group("/skills")
	skillController := controller.InitSkillController()
	m := middleware.InitMiddleware()

	skillRoute.GET("/", m.Cache(), skillController.Search())
	skillRoute.GET("/:id", m.Cache(), skillController.Get())

	skillRoute.Use(m.Authenticate())

	skillRoute.POST("/", skillController.Create())
	skillRoute.PUT("/:id", skillController.Update())
	skillRoute.DELETE("/:id", skillController.Delete())
}
