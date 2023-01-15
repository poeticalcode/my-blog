package service

import (
	"fmt"
	"github.com/he-wen-yao/my-blog/server/model/vo"
	"testing"
)

func TestArticleList(t *testing.T) {
	service := articleService{}
	list, _ := service.ArticleList(&vo.PagingParam{PageNum: 1, PageSize: 5})
	fmt.Println(list)
}
