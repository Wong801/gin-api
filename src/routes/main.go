package route

import (
	"net/http"

	"github.com/Wong801/gin-api/src/config"
	entity "github.com/Wong801/gin-api/src/entities"
	middleware "github.com/Wong801/gin-api/src/middlewares"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

type handler struct {
	router *gin.Engine
}

func InitRoutes() handler {
	r := handler{
		router: gin.New(),
	}

	store := cookie.NewStore([]byte(config.GetEnv("JWT_SECRET", "secret")))
	r.router.Use(sessions.Sessions("session", store))

	r.router.Use(csrf.Middleware(csrf.Options{
		Secret: config.GetEnv("CSRF_SECRET", "secret"),
		ErrorFunc: func(c *gin.Context) {
			c.AbortWithStatusJSON(http.StatusBadRequest, &entity.HttpResponse{
				Success: false,
				Data:    "CSRF token mismatch",
			})
		},
	}))

	m := middleware.InitMiddleware()

	r.router.StaticFS("/public", gin.Dir("public", false))
	r.router.Use(gin.Logger(), m.Recovery(), m.Response())

	root := r.router.Group("/")
	v1 := r.router.Group("/api/v1")
	r.addRoot(root)
	r.addUsers(v1)
	r.addCompanies(v1)

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
