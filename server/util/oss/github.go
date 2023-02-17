package oss

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path"
	"strings"

	"github.com/google/uuid"
	"github.com/he-wen-yao/my-blog/server/config"
	"github.com/he-wen-yao/my-blog/server/util"
)

type github struct{}

// UploadFile 上传文件
func (github) UploadFile(file io.Reader, fileName string) (url string, err error) {
	sourcebuffer, _ := io.ReadAll(file)
	github := config.GlobalConfig.GithubOSS
	dateYMD := strings.Split(util.Now(), " ")[0]
	// 使用 UUID 重新命名
	newFileName := uuid.New().String() + path.Ext(fileName)
	// 拼接上传的 API 地址
	uploadApi := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s/%s", github.UserName, github.Repository, dateYMD, newFileName)
	// 拼接 JSON 数据
	j, err := json.Marshal(map[string]interface{}{"message": fmt.Sprintf("%s:upload %s", github.Message, fileName), "content": base64.StdEncoding.EncodeToString(sourcebuffer)})
	if err != nil {
		return
	}
	payload := strings.NewReader(string(j))

	req, _ := http.NewRequest("PUT", uploadApi, payload)
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", github.Token))
	// 执行请求
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()
	res, err := util.ParseResponse(response)
	if err != nil {
		return
	}
	content, _ := res["content"].(map[string]interface{})
	switch response.StatusCode {
	// 更新资源
	case 200:
	case 201:
		url = fmt.Sprintf("%s/%s/%s@%s/%s", github.CDN, github.UserName, github.Repository, github.Branch, content["path"])
	default:
		url = ""
		err = errors.New("更新或创建失败")
		return
	}
	return
}
