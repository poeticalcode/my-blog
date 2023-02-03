package entity

import (
	"time"

	"gorm.io/gorm"
)

type baseModel struct {
	Id        int            `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"` // 更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"<-:false" `      // 删除时间
}
