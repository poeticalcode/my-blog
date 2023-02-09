package db

import (
	"github.com/he-wen-yao/my-blog/server/config"
	"github.com/he-wen-yao/my-blog/server/model/vo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	dsn := config.GlobalConfig.MySQL.Dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.Callback()
	if err != nil {
		panic(err)
	}
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
