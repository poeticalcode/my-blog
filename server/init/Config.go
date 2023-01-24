package init

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件发生错误：%s", err))
	}
	fmt.Sprintf("配置文件加载成功")
}
