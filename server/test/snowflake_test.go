package test

import (
	"github.com/he-wen-yao/my-blog/server/util/snowflake"
	"testing"
)

func Test_Snowflake(t *testing.T) {
	snowflakeId := snowflake.Snowflake.NextVal()

	t.Log(snowflakeId)
}
