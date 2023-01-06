package model

import (
	"time"
)

type UserBase struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	FirstName string    `json:"firstName" binding:"required"`
	LastName  string    `json:"lastName" binding:"required"`
	Username  string    `json:"username" binding:"required" gorm:"uniqueIndex"`
	Phone     string    `json:"phone" binding:"omitempty,e164"`
	Email     string    `json:"email" binding:"required,email" gorm:"uniqueIndex"`
	DoB       time.Time `json:"dob" binding:"required" time_format:"2006-01-02" gorm:"column:date_of_birth"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime:true"`
}

type User struct {
	UserBase
	Password string `json:"password" binding:"required,min=8"`
}

type UserLogin struct {
	Username string `json:"username" binding:"required_without=Email,omitempty"`
	Email    string `json:"email" binding:"required_without=Username,omitempty,email"`
	Password string `json:"password" binding:"required,omitempty"`
}

type UserChangePassword struct {
	OldPassword       string `json:"oldPassword" binding:"required"`
	VerifyOldPassword string `json:"verifyOldPassword" binding:"required"`
	NewPassword       string `json:"newPassword" binding:"required,min=8"`
}
