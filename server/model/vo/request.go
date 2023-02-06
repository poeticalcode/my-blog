package vo

import "errors"

// PagingParam 接收分页参数
type PagingParam struct {
	PageNum  int `form:"page_num"`
	PageSize int `form:"page_size"`
}

// 校验参数
func (parm *PagingParam) Check() error {
	if parm.PageNum <= 0 {
		return errors.New("PageNum 不能为空")
	}
	if parm.PageSize <= 0 {
		return errors.New("PageSize 不能为空")
	}
	return nil
}

// 获取偏移量
func (parm *PagingParam) Offset() (offset int) {
	offset = (parm.PageNum - 1) * parm.PageSize
	if offset < 0 {
		offset = 0
	}
	return
}

// 获取限制数量
func (parm *PagingParam) Limit() (limit int) {
	limit = parm.PageSize
	return
}
