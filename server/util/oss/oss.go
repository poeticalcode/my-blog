package oss

import "io"

var Github = new(github)

// Uploader 上传文件接口
type Uploader interface {
	Upload(reader io.Reader, fileName string) (url string, err error)
}

// Downloader 文件下载接口
type Downloader interface {
	Download(reader *io.Writer, path string) error
}
