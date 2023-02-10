package main

import (
	"fmt"

	"github.com/he-wen-yao/my-blog/server/config"
	"github.com/he-wen-yao/my-blog/server/db"
	init_ "github.com/he-wen-yao/my-blog/server/init"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	c := config.GlobalConfig
	db.DB = db.InitDB()
	app := init_.RouterInit()
	err := app.Run(fmt.Sprintf(":%s", c.Port))
	if err != nil {
		return
	}
}
