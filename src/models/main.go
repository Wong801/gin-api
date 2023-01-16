package model

import "time"

type Timestamp struct {
	CreatedAt time.Time `json:"createdAt" form:"createdAt" gorm:"type:timestamp;default:current_timestamp;autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt" gorm:"type:timestamp;default:current_timestamp;autoUpdateTime"`
}
