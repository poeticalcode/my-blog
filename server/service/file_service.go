package service

import (
	"errors"
	"fmt"
	"mime/multipart"
	"path"

	"github.com/he-wen-yao/my-blog/server/util/oss"
)

// 文件上传服务
type fileService struct{}

var allowExtMap = map[string]bool{
	".jpg":  true,
	".png":  true,
	".gif":  true,
	".jpeg": true,
}

func (fileService) UploadImage(fileHeader *multipart.FileHeader) (url string, err error) {
	// 获取扩展名
	extName := path.Ext(fileHeader.Filename)
	if _, ok := allowExtMap[extName]; !ok {
		err = errors.New("图片后缀名不合法")
		return
	}
	fileName := fileHeader.Filename + extName
	file, err := fileHeader.Open()
	if err != nil {
		err = errors.New(fmt.Sprintf("图片打开失,err = %s", err))
	}
	url, err = oss.Github.UploadFile(file, fileName)
	return
}
