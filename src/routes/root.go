package route

import (
	controller "github.com/Wong801/gin-api/src/controllers"
	"github.com/gin-gonic/gin"
)

func (r handler) addRoot(rg *gin.RouterGroup) {
	rc := controller.InitRootController()

	rg.GET("/stats", rc.GetStats())
	rg.GET("/ping", rc.Ping())
}
