package model

import "time"

type Timestamp struct {
	CreatedAt time.Time `json:"createdAt" form:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt" gorm:"autoUpdateTime:true"`
}
