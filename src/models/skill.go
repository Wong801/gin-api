package model

type Skill struct {
	ID   int    `json:"id" form:"id" gorm:"primaryKey"`
	Name string `json:"name" form:"name" binding:"required" gorm:"uniqueIndex"`
	Timestamp
}
