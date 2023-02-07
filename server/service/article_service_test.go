package service_test

import (
	"fmt"
	"testing"

	"github.com/he-wen-yao/my-blog/server/service"

	"github.com/he-wen-yao/my-blog/server/model/vo"
)

func TestArticleList(t *testing.T) {
	articleService := service.ArticleService
	list, _ := articleService.ArticleList(&vo.PagingParam{PageNum: 1, PageSize: 5})
	fmt.Println(list)
}

func TestArticleService_DeleteArticleById(t *testing.T) {

}

func TestArticleService_FetchArticleById(t *testing.T) {
	service := service.ArticleService
	article := service.FetchArticleById(106)
	fmt.Println(article)
}
