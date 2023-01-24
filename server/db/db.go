package db

import (
	"github.com/he-wen-yao/my-blog/server/config"
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
