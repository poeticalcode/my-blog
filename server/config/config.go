package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	"github.com/he-wen-yao/my-blog/server/util/path"
)

// 读取配置文件
func readConfig() error {
	path := path.GetProjectRootPath()
	GlobalConfig.ProjectRootPath = path
	fmt.Printf("path =========== %s", path)
	log.Printf("path =========== %s", path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(fmt.Sprintf("%s/config", path))
	return viper.ReadInConfig()
}

func init() {
	err := readConfig()
	if err != nil {
		panic(err)
	}

	GlobalConfig.Port = viper.GetString("server.port")
	GlobalConfig.MySQL.Host = viper.GetString("mysql.host")
	GlobalConfig.MySQL.Port = viper.GetString("mysql.port")
	GlobalConfig.MySQL.UserName = viper.GetString("mysql.userName")
	GlobalConfig.MySQL.Password = viper.GetString("mysql.password")
	GlobalConfig.MySQL.Charset = viper.GetString("mysql.charset")
	GlobalConfig.MySQL.DataBase = viper.GetString("mysql.database")
	// 初始化 githubOSS
	GlobalConfig.GithubOSS.Message = viper.GetString("githubOOS.message")
	GlobalConfig.GithubOSS.Token = viper.GetString("githubOOS.token")
	GlobalConfig.GithubOSS.UserName = viper.GetString("githubOOS.username")
	GlobalConfig.GithubOSS.Repository = viper.GetString("githubOOS.repository")
	GlobalConfig.GithubOSS.CDN = viper.GetString("githubOOS.cdn")
	GlobalConfig.GithubOSS.Branch = viper.GetString("githubOOS.branch")

}

var GlobalConfig = new(globalConfig)
