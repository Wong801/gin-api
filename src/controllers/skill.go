package controller

import (
	"net/http"

	"github.com/Wong801/gin-api/src/api"
	model "github.com/Wong801/gin-api/src/models"
	service "github.com/Wong801/gin-api/src/services"
	"github.com/gin-gonic/gin"
)

type skillUri struct {
	Id int `uri:"id" binding:"required"`
}

type SkillController struct {
	s service.CRUDService
}

func InitSkillController() *SkillController {
	return &SkillController{
		s: *service.InitCRUDService("skill", model.Skill{}),
	}
}

func (cc SkillController) Search() func(c *gin.Context) {
	return func(c *gin.Context) {
		status, list, errService := cc.s.Search([]model.Skill{})
		c.Set("status", status)
		if errService != nil {
			c.Set("error", api.MakeResultError(errService))
			return
		}

		c.Set("data", list)
		c.Next()
	}
}

func (cc SkillController) Get() func(c *gin.Context) {
	return func(c *gin.Context) {
		var cUri skillUri

		err := c.ShouldBindUri(&cUri)
		if err != nil {
			c.Set("status", http.StatusBadRequest)
			c.Set("error", api.MakeRequestError(err))
			return
		}

		status, data, errService := cc.s.Get(cUri.Id)
		c.Set("status", status)
		if errService != nil {
			c.Set("error", api.MakeResultError(errService))
			return
		}

		c.Set("data", data)
		c.Next()
	}
}

func (cc SkillController) Create() func(c *gin.Context) {
	return func(c *gin.Context) {
		var skill model.Skill

		err := c.ShouldBind(&skill)
		if err != nil {
			c.Set("status", http.StatusBadRequest)
			c.Set("error", api.MakeRequestError(err))
			return
		}
		status, data, errService := cc.s.Create(skill)
		c.Set("status", status)
		if errService != nil {
			c.Set("error", api.MakeResultError(errService))
			return
		}

		c.Set("data", data)
		c.Next()
	}
}

func (cc SkillController) Update() func(c *gin.Context) {
	return func(c *gin.Context) {
		var cUri skillUri
		var skill model.Skill

		c.ShouldBindUri(&cUri)
		err := c.ShouldBind(&skill)
		if err != nil {
			c.Set("status", http.StatusBadRequest)
			c.Set("error", err)
			return
		}

		status, data, errService := cc.s.Update(cUri.Id, skill)
		c.Set("status", status)
		if errService != nil {
			c.Set("error", api.MakeResultError(errService))
			return
		}

		c.Set("data", data)
		c.Next()
	}
}

func (cc SkillController) Delete() func(c *gin.Context) {
	return func(c *gin.Context) {
		var cUri skillUri

		err := c.ShouldBindUri(&cUri)
		if err != nil {
			c.Set("status", http.StatusBadRequest)
			c.Set("error", api.MakeRequestError(err))
			return
		}

		status, data, errService := cc.s.Delete(cUri.Id)
		c.Set("status", status)
		if errService != nil {
			c.Set("error", api.MakeResultError(errService))
			return
		}

		c.Set("data", data)
		c.Next()
	}
}
