package service

import (
	"github.com/he-wen-yao/my-blog/server/db"
	"github.com/he-wen-yao/my-blog/server/model/do"
)

type articleService struct{}

// FetchArticleById 通过 ID 获取文章
func (articleService) FetchArticleById(id int64) *do.Article {
	var article *do.Article
	db.DB().First(article, "id = ?", id)
	return article
}

func (articleService) CreateArticle(article *do.Article) bool {
	res := db.DB().Create(article).RowsAffected
	return res != 0
}
