package test

import (
	"testing"

	"github.com/he-wen-yao/my-blog/server/util/snowflake"
)

func Test_Snowflake(t *testing.T) {
	snow := snowflake.NewSnowflake(0, 0)
	snowflakeId := snow.NextVal()
	t.Log(snowflakeId)
}
