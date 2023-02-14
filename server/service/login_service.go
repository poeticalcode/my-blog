package service

import (
	"github.com/he-wen-yao/my-blog/server/db"
	"github.com/he-wen-yao/my-blog/server/model/entity"
	"github.com/he-wen-yao/my-blog/server/model/req"
)

type loginService struct{}

// Login 登录业务实现
func (loginService) Login(loginReq *req.LoginReqParam) (user *entity.User) {
	db.DB.Where("email = ? and password = ? ", loginReq.Email, loginReq.Password).First(&user)
	return
}
