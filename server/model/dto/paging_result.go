package dto

import "github.com/he-wen-yao/my-blog/server/model/vo"

// PagingResult 分页结果实体
type pagingResult struct {
	Total       int         `json:"total"`
	List        interface{} `json:"list"`
	CurrentPage int         `json:"current_page"`
	LastPage    int         `json:"last_page"`
	NextPage    int         `json:"next_page"`
	TotalPage   int         `json:"total_page"`
	PageSize    int         `json:"page_size"`
}

func GenPagingResult(list interface{}, param vo.PagingParam, total int) pagingResult {
	return pagingResult{
		Total:       total,
		List:        list,
		CurrentPage: param.PageNum,
		PageSize:    param.PageSize,
		NextPage:    param.PageNum + 1,
		LastPage:    param.PageNum - 1,
		TotalPage:   total/param.PageSize + 1,
	}
}
