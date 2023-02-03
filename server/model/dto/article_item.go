package dto

import (
	"time"
)

type ArticleItem struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at" ` // 创建时间
	UpdatedAt   time.Time `json:"updated_at" ` // 更新时间
}
