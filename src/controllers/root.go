package controller

import (
	"net/http"

	"github.com/Wong801/gin-api/src/api"
	"github.com/Wong801/gin-api/src/db"
	service "github.com/Wong801/gin-api/src/services"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

type RootController struct {
	s service.RootService
}

func InitRootController() *RootController {
	rc := &RootController{
		s: *service.InitRootService(),
	}
	return rc
}

func (rc RootController) GetStats() func(c *gin.Context) {
	return func(c *gin.Context) {
		status, stats, err := rc.s.GetStats(db.InitDB())
		c.Set("status", status)
		if err != nil {
			c.Set("error", api.MakeResultError(err))
			return
		}
		c.Set("data", stats)

		c.Next()
	}
}

func (rc RootController) Ping() func(c *gin.Context) {
	return func(c *gin.Context) {
		status, message := rc.s.Ping()
		c.Set("status", status)
		c.Set("data", message)

		c.Next()
	}
}

func (rc RootController) GetToken() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Set("status", http.StatusOK)
		c.Set("data", csrf.GetToken(c))

		c.Next()
	}
}
