package config

import (
	"github.com/spf13/viper"
)

var GlobalConfig = new(globalConfig)

// 读取配置文件
func readConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	return viper.ReadInConfig()
}

func init() {
	err := readConfig()
	if err != nil {
		return
	}
	GlobalConfig.Port = viper.GetString("server.port")
}
