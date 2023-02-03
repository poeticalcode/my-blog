package vo

// PagingParam 接收分页参数
type PagingParam struct {
	PageNum  int `form:"page_num"`
	PageSize int `form:"page_size"`
}
