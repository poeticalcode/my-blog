package main

import (
	"fmt"
	"github.com/he-wen-yao/my-blog/server/config"
	"github.com/he-wen-yao/my-blog/server/db"
	"github.com/he-wen-yao/my-blog/server/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	c := config.GlobalConfig
	db.DB = db.InitDB()
	app := initialize.RouterInit()
	err := app.Run(fmt.Sprintf(":%s", c.Port))
	if err != nil {
		return
	}
}
