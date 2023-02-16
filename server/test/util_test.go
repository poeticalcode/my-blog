package test

import (
	"fmt"
	"testing"

	"github.com/he-wen-yao/my-blog/server/util"
	"github.com/he-wen-yao/my-blog/server/util/path"
	"github.com/he-wen-yao/my-blog/server/util/oos"
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

func TestPath_GetCurrentAbPath(t *testing.T) {
	path := path.GetProjectRootPath()
	fmt.Println(path)
}

func TestUploadGithub_Upload(t *testing.T) {
	err := oos.Github.UploadFileToGithub()
	if err != nil {
		t.Error(err)
	}
	t.Log()
}
