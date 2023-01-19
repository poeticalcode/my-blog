package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/he-wen-yao/my-blog/server/model/do"
	"github.com/he-wen-yao/my-blog/server/model/dto"
	"github.com/he-wen-yao/my-blog/server/model/vo"
	"github.com/he-wen-yao/my-blog/server/service"
	"strconv"
	"strings"
)

// ArticleApi 文章相关 api 集合
type ArticleApi struct{}

// CreateArticle 添加文章
func (ArticleApi) CreateArticle(c *gin.Context) {
	var article do.Article
	err := c.ShouldBindJSON(&article)
	if err != nil {
		dto.ResponseGen.FailWithMessage("提交的格式有误", c)
		return
	}
	if len(strings.Trim(article.Title, "")) == 0 {
		dto.ResponseGen.FailWithMessage("标题不能为空", c)
		return
	}
	if len(strings.Trim(article.MdText, "")) == 0 {
		dto.ResponseGen.FailWithMessage("内容不你能为空", c)
		return
	}

	if res := service.ArticleService.CreateArticle(&article); !res {
		dto.ResponseGen.FailWithMessage("插入失败", c)
		return
	}
	dto.ResponseGen.OkWithDetailed(article, "插入成功", c)
}

func (ArticleApi) UpdateArticle(c *gin.Context) {
	var article do.Article
	err := c.ShouldBindJSON(&article)
	if err != nil {
		dto.ResponseGen.FailWithMessage("提交的格式有误", c)
		return
	}
	if len(strings.Trim(article.Title, "")) == 0 {
		dto.ResponseGen.FailWithMessage("标题不能为空", c)
		return
	}
	if len(strings.Trim(article.MdText, "")) == 0 {
		dto.ResponseGen.FailWithMessage("内容不你能为空", c)
		return
	}

	if res, _ := service.ArticleService.UpdateArticle(&article); !res {
		dto.ResponseGen.FailWithMessage("更新失败", c)
		return
	}
	dto.ResponseGen.OkWithDetailed(article, "更新成功", c)
}

// FetchArticleDetail 获取文章详情
func (ArticleApi) FetchArticleDetail(c *gin.Context) {
	id := c.Param("id")
	if id, err := strconv.ParseInt(id, 10, 64); err == nil {
		article := service.ArticleService.FetchArticleById(id)
		dto.ResponseGen.OkWithDetailed(article, "获取成功", c)
		return
	}
	dto.ResponseGen.FailWithMessage("文章 ID 填写有误", c)
}

// FetchArticleListByPaging 分页获取文章列表
func (ArticleApi) FetchArticleListByPaging(c *gin.Context) {
	param := vo.PagingParam{}
	err := c.ShouldBindJSON(&param)
	if err != nil {
		dto.ResponseGen.FailWithMessage("提交的格式有误", c)
		return
	}
	list, _ := service.ArticleService.ArticleList(&param)
	dto.ResponseGen.OkWithData(list, c)
}
