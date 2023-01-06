package route

import (
	"net/http"

	entity "github.com/Wong801/gin-api/src/entities"
	middleware "github.com/Wong801/gin-api/src/middlewares"
	"github.com/gin-gonic/gin"
)

type handler struct {
	router *gin.Engine
}

func InitRoutes() handler {
	r := handler{
		router: gin.New(),
	}
	m := middleware.InitMiddleware()

	r.router.Use(gin.Logger(), m.Recovery(), m.Response())

	v1 := r.router.Group("/api/v1")

	r.addUsers(v1)

	return r
}

func (r handler) Run(addr string) error {
	r.router.NoRoute(func(c *gin.Context) {
		c.Set("status", http.StatusNotFound)
		c.Set("error", entity.ResultError{Reason: "Not Found"})
		c.Next()
	})
	r.router.NoMethod(func(c *gin.Context) {
		c.Set("status", http.StatusMethodNotAllowed)
		c.Set("error", entity.ResultError{Reason: "Method Not Allowed"})
		c.Next()
	})
	return r.router.Run(addr)
}
