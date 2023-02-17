package oss

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/he-wen-yao/my-blog/server/config"
	"github.com/he-wen-yao/my-blog/server/util"
)

type github struct{}

// UploadFile 上传文件
func (github) UploadFile(file io.Reader, fileName string) (path string, err error) {
	sourcebuffer, _ := io.ReadAll(file)
	github := config.GlobalConfig.GithubOSS
	dateYMD := strings.Split(util.Now(), " ")[0]
	path = fmt.Sprintf("%s/%s/contents/%s/%s", github.UserName, github.Repository, dateYMD, fileName)
	uploadApi := fmt.Sprintf("https://api.github.com/repos/%s", path)
	j, err := json.Marshal(map[string]interface{}{"message": github.Message, "content": base64.StdEncoding.EncodeToString(sourcebuffer)})
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
	_, err = util.ParseResponse(response)
	if err != nil {
		return
	}
	switch response.StatusCode {
	// 更新资源
	case 200:
	case 201:
	default:
		path = ""
		err = errors.New("更新或创建失败")
		return
	}
	return
}
