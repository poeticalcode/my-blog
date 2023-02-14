package test

import (
	"testing"

	"github.com/he-wen-yao/my-blog/server/util"
)

func TestJwt_GenerateToken(t *testing.T) {
	token, err := util.GenerateToken("he.wenyao", "he.123465")
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}

func TestJwt_ParseToken(t *testing.T) {
	token, _ := util.GenerateToken("he.wenyao", "he.123456")
	claims, err := util.ParseToken(token)
	if err != nil {
		t.Error(err)
	}
	t.Log(claims)
	if claims.Name == "he.wenyao" && claims.Phone == "he.123456" {
		t.Log("Yes")
		return
	}
	t.Error("解析错误")
}
