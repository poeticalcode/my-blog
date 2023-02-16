package test

import (
	"fmt"
	"testing"

	"github.com/he-wen-yao/my-blog/server/config"
)

func TestMysqlConfig_Dsn(t *testing.T) {
	dsn := config.GlobalConfig.MySQL.Dsn()
	fmt.Println(dsn)
}
