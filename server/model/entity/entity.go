package entity

import (
	"log"

	"github.com/he-wen-yao/my-blog/server/global"
	"gorm.io/gorm"
)

// BaseModel 基础模型 [公共部分]
type BaseModel struct {
	ID        DbUInt64       `json:"id" gorm:"primaryKey;autoIncrement=false"`
	CreatedAt XTime          `json:"created_at"` // 创建时间
	UpdatedAt XTime          `json:"updated_at"` // 更新时间
	DeletedAt gorm.DeletedAt `json:"-"`          // 删除时间
}

// Article 文章模型
type Article struct {
	BaseModel
	Title       string `json:"title"`       // 文章标题
	Status      uint16 `json:"status"`      // 文章状态
	ViewNum     uint64 `json:"view_num"`    // 阅读数量
	Cover       string `json:"cover"`       // 文章封面
	MdText      string `json:"md_text"`     // markdown 文本
	Description string `json:"description"` // 文章描述
}

func (a *Article) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = DbUInt64(global.Snowflake.NextVal())
	log.Println(a)
	return
}

// TableName 返回 Article 模型对应的数据库表名
func (a *Article) TableName() string {
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

// User 用户模型
type User struct {
	BaseModel
	Email    string `json:"email"`     // 邮箱
	Password string `json:"password"`  // 密码
	NickName string `json:"nick_name"` // 昵称
}

// TableName 返回 User 模型对应的数据库表名
func (User) TableName() string {
	return "tb_user"
}
