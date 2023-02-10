package db

import (
	"github.com/he-wen-yao/my-blog/server/config"
	"github.com/he-wen-yao/my-blog/server/model/vo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := config.GlobalConfig.MySQL.Dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db_, _ := db.DB()
	if err != nil {
		db_.Close()
		panic(err)
	}
	// 设置连接池，空闲连接
	db_.SetMaxIdleConns(50)
	// 打开链接
	db_.SetMaxOpenConns(100)
	return db
}

// Paginate 分页查询公共部分
func Paginate(p *vo.PagingParam) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (p.PageNum - 1) * p.PageSize
		limit := p.PageSize
		return db.Offset(offset).Limit(limit)
	}
}
