package model

import (
	"time"
)

type UserBase struct {
	Id        int        `json:"id" form:"id" gorm:"primaryKey"`
	FirstName string     `json:"firstName" form:"firstName" binding:"required"`
	LastName  string     `json:"lastName" form:"lastName" binding:"required"`
	Username  string     `json:"username" form:"username" binding:"required" gorm:"uniqueIndex"`
	Phone     string     `json:"phone" form:"phone" binding:"omitempty,e164"`
	Email     string     `json:"email" form:"email" binding:"required,email" gorm:"uniqueIndex"`
	DoB       *time.Time `json:"dob" form:"dob" binding:"required" time_format:"2006-01-02" gorm:"column:date_of_birth"`
	CreatedAt time.Time  `json:"createdAt" form:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time  `json:"updatedAt" form:"updatedAt" gorm:"autoUpdateTime:true"`
}

type User struct {
	UserBase
	Password string `json:"password" form:"password" binding:"required,min=8"`
}

type UserLogin struct {
	Username string `json:"username" form:"username" binding:"required_without=Email,omitempty"`
	Email    string `json:"email" form:"email" binding:"required_without=Username,omitempty,email"`
	Password string `json:"password" form:"password" binding:"required,omitempty"`
}

type UserChangePassword struct {
	OldPassword       string `json:"oldPassword" form:"oldPassword" binding:"required"`
	VerifyOldPassword string `json:"verifyOldPassword" form:"verifyOldPassword" binding:"required"`
	NewPassword       string `json:"newPassword" form:"newPassword" binding:"required,min=8"`
}
