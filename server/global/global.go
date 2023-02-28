package global

import (
	"github.com/he-wen-yao/my-blog/server/util/snowflake"
	"gorm.io/gorm"
)

// 数据库实例
var DB *gorm.DB

// 雪花算法实例
var Snowflake *snowflake.Snowflake
