package entity

import (
	"time"

	"gorm.io/gorm"
)

// 基础模型 [公共部分]
type BaseModel struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"` // 创建时间
	UpdatedAt time.Time      `json:"updated_at"` // 更新时间
	DeletedAt gorm.DeletedAt `json:"-"`          // 删除时间
}

// 文章模型
type Article struct {
	BaseModel
	Title       string `json:"title"`       // 文章标题
	Cover       string `json:"cover"`       // 文章封面
	MdText      string `json:"md_text"`     // markdowm 文本
	Description string `json:"description"` // 文章描述
}

func (Article) TableName() string {
	return "tb_article"
}

// 标签模型
type Tag struct {
	BaseModel
	Name string `json:"name"` // 标签名称
}

func (Tag) TableName() string {
	return "tb_tag"
}
