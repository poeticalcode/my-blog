package req

import (
	"errors"
)

// PagingParam 接收分页参数
type PagingParam struct {
	PageNum  int `form:"page_num"`
	PageSize int `form:"page_size"`
}

// Check 校验参数
func (parm *PagingParam) Check() error {
	if parm.PageNum <= 0 {
		return errors.New("page_num 不能为空或者为0")
	}
	if parm.PageSize <= 0 {
		return errors.New("page_size 不能为空或者为0")
	}
	return nil
}

// Offset 获取偏移量
func (parm *PagingParam) Offset() (offset int) {
	offset = (parm.PageNum - 1) * parm.PageSize
	if offset < 0 {
		offset = 0
	}
	return
}

// Limit 获取限制数量
func (parm *PagingParam) Limit() (limit int) {
	limit = parm.PageSize
	return
}
