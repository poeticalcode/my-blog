package v1

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/he-wen-yao/my-blog/server/model/req"
	"github.com/he-wen-yao/my-blog/server/model/res"
	"github.com/he-wen-yao/my-blog/server/service"
	"github.com/he-wen-yao/my-blog/server/util"
)

type LoginApi struct{}

// Login 登录接口实现
func (LoginApi) Login(c *gin.Context) {
	var params req.LoginReqParam
	err := c.ShouldBindJSON(&params)
	if err != nil {
		res.ResponseGen.FailWithMessage("提交的格式有误", c)
		return
	}
	user := service.LoginService.Login(&params)
	log.Print(user)
	token, err := util.GenerateToken(user.Email, "he.123456")
	if err != nil {
		res.ResponseGen.FailWithMessage("JWT 生成 Token 失败", c)
	}
	res.ResponseGen.OkWithDetailed(map[string]interface{}{"token": token}, "登录成功", c)
}
