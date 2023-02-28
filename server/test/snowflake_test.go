package test

import (
	"testing"

	"github.com/he-wen-yao/my-blog/server/db"
	"github.com/he-wen-yao/my-blog/server/global"
	"github.com/he-wen-yao/my-blog/server/util/snowflake"
)

func Test_Snowflake(t *testing.T) {
	snow := snowflake.NewSnowflake(0, 0)
	snowflakeId := snow.NextVal()
	t.Log(snowflakeId)
}

func Test_FlushArticleId(t *testing.T) {
	snow := snowflake.NewSnowflake(0, 0)
	result := []int64{}
	global.DB = db.InitDB()
	global.DB.Table("tb_article").Select("id").Find(&result)
	for _, id := range result {
		snowflakeId := snow.NextVal()
		global.DB.Table("tb_article").Where("id = ?", id).UpdateColumn("id", snowflakeId)
		t.Log(snowflakeId)
	}
}
