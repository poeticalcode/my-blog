package config

import (
	"fmt"
	"testing"
)

func TestMysqlConfig_Dsn(t *testing.T) {
	dsn := GlobalConfig.MySQL.Dsn()
	fmt.Println(dsn)
}
