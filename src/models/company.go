package model

type Company struct {
	ID       int    `json:"id" form:"id" gorm:"primaryKey"`
	Name     string `json:"name" form:"name" binding:"required"`
	Logo     string `json:"logo" form:"logo"`
	Link     string `json:"link" form:"link" binding:"omitempty"`
	State    string `json:"state" form:"state" binding:"required"`
	City     string `json:"city" form:"city" binding:"required"`
	Province string `json:"province" form:"province" binding:"required"`
	Timestamp
}
