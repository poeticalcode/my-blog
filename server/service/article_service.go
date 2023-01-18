package service

import (
	"github.com/he-wen-yao/my-blog/server/db"
	"github.com/he-wen-yao/my-blog/server/model/do"
	"github.com/he-wen-yao/my-blog/server/model/vo"
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

func (articleService) ArticleList(param *vo.PagingParam) ([]do.Article, error) {
	pageNum := param.PageNum
	pageSize := param.PageSize
	articleList := make([]do.Article, 0)
	result := db.DB().Offset(pageNum).Limit(pageSize).Find(&articleList)
	if result.Error != nil {
		return nil, result.Error
	}
	return articleList, nil
}

func (articleService) UpdateArticle(article *do.Article) (bool, error) {
	res := db.DB().Updates(article)
	if res.Error != nil {
		return false, res.Error
	}
	return res.RowsAffected != 0, nil
}
