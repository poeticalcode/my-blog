package vo

// PagingParam 接收分页参数
type PagingParam struct {
	PageNum  int `json:"page_num"`
	PageSize int `json:"page_size"`
}
