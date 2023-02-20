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

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	c := config.GlobalConfig
	db.DB = db.InitDB()
	app := initialize.RouterInit()
	err := app.Run(fmt.Sprintf(":%s", c.Port))
	if err != nil {
		return
	}
}
