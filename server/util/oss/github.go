package oss

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/he-wen-yao/my-blog/server/config"
)

type github struct{}

// UploadFile 上传文件
func (github) UploadFile(file io.Reader, fileName string) error {
	imgByte, _ := io.ReadAll(file)
	// 取图片类型
	mimeType := http.DetectContentType(imgByte)
	imgBase64 := ""
	switch mimeType {
	case "image/jpeg":
		imgBase64 = "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(imgByte)
	case "image/png":
		imgBase64 = "data:image/png;base64," + base64.StdEncoding.EncodeToString(imgByte)
	}
	github := config.GlobalConfig.GithubOOS
	const LAYOUT = "2006-01-02 15:04:05"
	// 获取当前日期
	now := time.Now()
	now_ := strings.Split(now.Format(LAYOUT), " ")[0] // 获取 年 月 日
	// 2022-08-31 09:48:40
	uploadApi := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s/%s", github.UserName, github.Repository, now_, fileName)
	log.Println(uploadApi)
	p := map[string]string{"message": github.Message, "content": imgBase64}
	j, _ := json.Marshal(p)
	payload := strings.NewReader(string(j))
	log.Println(payload)
	req, _ := http.NewRequest("PUT", uploadApi, payload)
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", github.Token))
	// 执行请求
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	log.Println(response.)
	return nil
}
