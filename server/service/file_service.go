package service

import (
	"errors"
	"github.com/he-wen-yao/my-blog/server/util/oss"
	"mime/multipart"
	"path"
)

// 文件上传服务
type fileService struct{}

var allowExtMap = map[string]bool{
	".jpg":  true,
	".png":  true,
	".gif":  true,
	".jpeg": true,
}

func (fileService) uploadImage(fileHeader multipart.FileHeader) error {
	// 获取扩展名
	extName := path.Ext(fileHeader.Filename)
	if _, ok := allowExtMap[extName]; !ok {
		return errors.New("图片后缀名不合法")
	}
	fileName := fileHeader.Filename + extName
	file, _ := fileHeader.Open()
	oss.Github.UploadFile(file, fileName)
	return nil
}
