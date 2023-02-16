package oos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/he-wen-yao/my-blog/server/config"
)

type github struct{}

// 上传文件
func (github) UploadFile() error {
	fileName := "test.txt"
	github := config.GlobalConfig.GithubOOS

	const LAYOUT = "2006-01-02 15:04:05"
	// 获取当前日期
	now := time.Now()
	now_ := strings.Split(now.Format(LAYOUT), " ")[0] // 获取 年 月 日
	// 2022-08-31 09:48:40
	uploadApi := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s/%s", github.UserName, github.Repository, now_, fileName)
	p := map[string]string{"message": github.Message, "content": "bXkgbmV3IGZpbGUgY29udGVudHM="}
	j, _ := json.Marshal(p)
	payload := strings.NewReader(string(j))
	req, _ := http.NewRequest("PUT", uploadApi, payload)
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", github.Token))
	// 执行请求
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return nil
}
