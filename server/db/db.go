package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	dsn := "root:he..123456@tcp(150.158.95.91:3306)/self_blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.Callback()
	if err != nil {
		panic(err)
	}
	return db
}
