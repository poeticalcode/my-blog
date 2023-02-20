package test

import (
	"encoding/base64"
	"os"
	"testing"

	"github.com/he-wen-yao/my-blog/server/util/oss"
)

func TestUploadGithub(t *testing.T) {
	//获取本地文件
	file, err := os.Open("./test.png")
	if err != nil {
		t.Error(err)
		return
	}
	defer file.Close()
	path, err := oss.Github.Upload(file, "12222ddd2222dasdas22ddd222asda111.jgp")
	if err != nil {
		t.Errorf("上传失败 %s", err)
		return
	}
	t.Logf("上传成功 %s", path)
}

func TestBase64(t *testing.T) {
	tt, _ := base64.StdEncoding.Strict().DecodeString("bXkgbmV3IGZpbGUgY29udGVudHM=")
	t.Log(string(tt))

	t.Log(base64.StdEncoding.Strict().EncodeToString([]byte("my new file contents")))
}
