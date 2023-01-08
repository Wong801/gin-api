package route

import (
	controller "github.com/Wong801/gin-api/src/controllers"
	middleware "github.com/Wong801/gin-api/src/middlewares"
	"github.com/gin-gonic/gin"
)

func (h handler) addCompanies(rg *gin.RouterGroup) {
	companyRoute := rg.Group("/companies")
	companyController := controller.InitCompanyController()
	m := middleware.InitMiddleware()

	companyRoute.GET("/", m.Cache(), companyController.Search())
	companyRoute.GET("/:id", m.Cache(), companyController.Get())

	companyRoute.Use(m.Authenticate())

	companyRoute.POST("/", m.LogoHandler("logoImg", "public/assets/images/company_logo/"), companyController.Create())
	companyRoute.PUT("/:id", m.LogoHandler("logoImg", "public/assets/images/company_logo/"), companyController.Update())
	companyRoute.DELETE("/:id", companyController.Delete())
}
