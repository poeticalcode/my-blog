package entity

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel 基础模型 [公共部分]
type BaseModel struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"` // 创建时间
	UpdatedAt time.Time      `json:"updated_at"` // 更新时间
	DeletedAt gorm.DeletedAt `json:"-"`          // 删除时间
}

// Article 文章模型
type Article struct {
	BaseModel
	Title       string `json:"title"`       // 文章标题
	Status      uint16 `json:"status"`      // 文章状态
	ViewNum     uint64 `json:"view_num"`    // 
	Cover       string `json:"cover"`       // 文章封面
	MdText      string `json:"md_text"`     // markdowm 文本
	Description string `json:"description"` // 文章描述
}

// TableName 返回 Article 模型对应的数据库表名
func (Article) TableName() string {
	return "tb_article"
}

// Tag 标签模型
type Tag struct {
	BaseModel
	Name string `json:"name"` // 标签名称
}

// TableName 返回 Tag 模型对应的数据库表名
func (Tag) TableName() string {
	return "tb_tag"
}
