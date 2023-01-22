package main

import (
	init_ "github.com/he-wen-yao/my-blog/server/init"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {

	app := init_.RouterInit()
	err := app.Run(":8080")
	if err != nil {
		return
	}
}
