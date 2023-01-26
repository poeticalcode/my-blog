package config

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestReadConfig(t *testing.T) {
	err := readConfig()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(viper.Get("server.port"))
}
