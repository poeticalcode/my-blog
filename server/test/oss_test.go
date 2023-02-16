package test

import (
	"github.com/he-wen-yao/my-blog/server/util/oss"
	"os"
	"testing"
)

func TestUploadGithub(t *testing.T) {
	//获取本地文件
	file, err := os.Open("./test.png")
	if err != nil {
		t.Error(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	err = oss.Github.UploadFile(file, "test.jgp")
	if err != nil {
		t.Error("上传失败")
		return
	}
	t.Log("上传成功")
}
