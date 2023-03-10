package service

import (
	"fmt"
	"log"

	"github.com/he-wen-yao/my-blog/server/db"

	"github.com/he-wen-yao/my-blog/server/model/entity"

	"github.com/he-wen-yao/my-blog/server/model/req"
)

type articleService struct{}

// FetchArticleById 通过 ID 获取文章
func (articleService) FetchArticleById(id int64) *entity.Article {
	var article entity.Article
	log.Printf("id = %d", id)
	res := db.DB.First(&article, id)
	if res.RowsAffected == 0 {
		return nil
	}
	return &article
}

// CreateArticle 创建文章
func (articleService) CreateArticle(article *entity.Article) bool {
	res := db.DB.Create(article).RowsAffected
	return res != 0
}

// ArticleList 分页获取文章信息
func (articleService) ArticleList(param *req.PagingParam) ([]entity.Article, error) {
	articleList := make([]entity.Article, 0)
	// 除了 md_text 字段之外都需要保留
	result := db.DB.Omit("md_text").Order("created_at desc").Scopes(db.Paginate(param)).Find(&articleList)
	if result.Error != nil {
		return nil, result.Error
	}
	return articleList, nil
}

// FetchTotalNum  获取文章总数
func (articleService) FetchTotalNum() (res int64) {
	db.DB.Model(entity.Article{}).Count(&res)
	return res
}

// UpdateArticle 更新文章
func (articleService) UpdateArticle(article *entity.Article) (bool, error) {
	res := db.DB.Updates(article)
	if res.Error != nil {
		return false, res.Error
	}
	return res.RowsAffected != 0, nil
}

// DeleteArticleById 删除文章
func (articleService) DeleteArticleById(id int64) (bool, error) {
	// 根据主键删除
	res := db.DB.Delete(&entity.Article{}, id)
	if res.Error != nil {
		return false, res.Error
	}
	return res.RowsAffected != 0, nil
}

// SwitchArticleStatusById  切换文章状态
func (articleService) SwitchArticleStatusById(id int64) (bool, error) {
	// SwitchArticleStatusById
	// todo
	return false, nil
}

// StickyArticleById  置顶文章
func (articleService) StickyArticleById(id int64) (bool, error) {
	// 根据主键置顶文章
	// todo
	return false, nil
}

// GetRecordsCountByCreateAt 统计每天发布文章数量
func (articleService) GetRecordsCountByCreateAt() ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	groupBy := "DATE_FORMAT(created_at,\"%Y-%m-%d\")"
	article := entity.Article{}
	err := db.DB.Table(article.TableName()).
		Select(fmt.Sprintf("%s as date", groupBy), "count(*) as count").
		Group(groupBy).
		Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
