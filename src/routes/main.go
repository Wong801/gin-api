package route

import (
	"net/http"

	entity "github.com/Wong801/gin-api/src/entities"
	middleware "github.com/Wong801/gin-api/src/middlewares"
	"github.com/gin-gonic/gin"
)

type route struct {
	router *gin.Engine
}

func InitRoutes() route {
	r := route{
		router: gin.New(),
	}

	r.router.Use(gin.Logger(), middleware.Recovery())

	v1 := r.router.Group("/api/v1")

	r.addUsers(v1)

	r.router.Use(middleware.Response())

	return r
}

func (r route) Run(addr string) error {
	r.router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, &entity.HttpResponse{
			Success: false,
			Data: map[string]string{
				"message": "Not Found",
			},
		})
	})

	r.router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, &entity.HttpResponse{
			Success: false,
			Data: map[string]string{
				"message": "Method Not Allowed",
			},
		})
	})

	return r.router.Run(addr)
}
