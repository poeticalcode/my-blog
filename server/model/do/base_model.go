package do

import (
	"gorm.io/gorm"
	"time"
)

type baseModel struct {
	id        gorm.DeletedAt
	Id        int            `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"` // 更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"<-:false" `      // 删除时间
}
