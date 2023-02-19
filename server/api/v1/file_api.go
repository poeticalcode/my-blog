package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/he-wen-yao/my-blog/server/model/res"
	"github.com/he-wen-yao/my-blog/server/service"
)

type FileApi struct{}

// UploadImage Login 登录接口实现
func (FileApi) UploadImage(c *gin.Context) {
	// 单文件
	file, err := c.FormFile("file")
	if err != nil {
		res.ResponseGen.FailWithMessage(fmt.Sprintf("获取文件错误 %s", err), c)
		return
	}
	url, err := service.FileService.UploadImage(file)
	if err != nil {
		res.ResponseGen.FailWithMessage(fmt.Sprintf("上传失败 %s", err), c)
		return
	}
	res.ResponseGen.OkWithDetailed(map[string]interface{}{"url": url}, "上传成功", c)
}
