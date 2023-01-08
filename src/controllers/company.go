package controller

import (
	"net/http"

	model "github.com/Wong801/gin-api/src/models"
	service "github.com/Wong801/gin-api/src/services"
	"github.com/gin-gonic/gin"
)

type companyQuery struct {
	Name string `json:"name" form:"name"`
}

type companyUri struct {
	Id int `uri:"id" binding:"required"`
}

type CompanyController struct {
	s service.CompanyService
}

func InitCompanyController() *CompanyController {
	return &CompanyController{
		s: *service.InitCompanyService(),
	}
}

func (cc CompanyController) Search() func(c *gin.Context) {
	return func(c *gin.Context) {
		var cQuery companyQuery

		err := c.ShouldBind(&cQuery)
		if err != nil {
			cQuery.Name = ""
		}

		status, list, errService := cc.s.Search(cQuery.Name)
		c.Set("status", status)
		if errService != nil {
			c.Set("error", errService)
			return
		}

		c.Set("data", list)
		c.Next()
	}
}

func (cc CompanyController) Get() func(c *gin.Context) {
	return func(c *gin.Context) {
		var cUri companyUri

		err := c.ShouldBindUri(&cUri)
		if err != nil {
			c.Set("status", http.StatusBadRequest)
			c.Set("error", err)
		}

		status, data, errService := cc.s.Get(cUri.Id)
		c.Set("status", status)
		if errService != nil {
			c.Set("error", errService)
			return
		}

		c.Set("data", data)
		c.Next()
	}
}

func (cc CompanyController) Create() func(c *gin.Context) {
	return func(c *gin.Context) {
		var company model.Company

		err := c.ShouldBind(&company)
		if err != nil {
			c.Set("status", http.StatusBadRequest)
			c.Set("error", err)
		}
		if logoPath, ok := c.Get("company_logo"); ok {
			company.Logo = logoPath.(string)
		}

		status, data, errService := cc.s.Create(company)
		c.Set("status", status)
		if errService != nil {
			c.Set("error", errService)
			return
		}

		c.Set("data", data)
		c.Next()
	}
}

func (cc CompanyController) Update() func(c *gin.Context) {
	return func(c *gin.Context) {
		var cUri companyUri
		var company model.Company

		c.ShouldBindUri(&cUri)
		err := c.ShouldBind(&company)
		if err != nil {
			c.Set("status", http.StatusBadRequest)
			c.Set("error", err)
		}
		if logoPath, ok := c.Get("company_logo"); ok {
			company.Logo = logoPath.(string)
		}

		status, data, errService := cc.s.Update(cUri.Id, company)
		c.Set("status", status)
		if errService != nil {
			c.Set("error", errService)
			return
		}

		c.Set("data", data)
		c.Next()
	}
}

func (cc CompanyController) Delete() func(c *gin.Context) {
	return func(c *gin.Context) {
		var cUri companyUri

		err := c.ShouldBindUri(&cUri)
		if err != nil {
			c.Set("status", http.StatusBadRequest)
			c.Set("error", err)
		}

		status, data, errService := cc.s.Delete(cUri.Id)
		c.Set("status", status)
		if errService != nil {
			c.Set("error", errService)
			return
		}

		c.Set("data", data)
		c.Next()
	}
}
