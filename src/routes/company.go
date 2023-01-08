package route

import (
	controller "github.com/Wong801/gin-api/src/controllers"
	middleware "github.com/Wong801/gin-api/src/middlewares"
	"github.com/gin-gonic/gin"
)

func logoHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		file, err := c.FormFile("logoFile")
		if err == nil {
			dst := "public/assets/images/company_logo/" + file.Filename

			c.Set("company_logo", "/"+dst)
			c.SaveUploadedFile(file, dst)
		}
		c.Next()
	}
}

func (h handler) addCompanies(rg *gin.RouterGroup) {
	companyRoute := rg.Group("/companies")
	companyController := controller.InitCompanyController()
	m := middleware.InitMiddleware()

	companyRoute.GET("/", m.Cache(), companyController.Search())
	companyRoute.GET("/:id", m.Cache(), companyController.Get())

	companyRoute.Use(m.Authenticate())

	companyRoute.POST("/", logoHandler(), companyController.Create())
	companyRoute.PUT("/:id", logoHandler(), companyController.Update())
	companyRoute.DELETE("/:id", companyController.Delete())
}
