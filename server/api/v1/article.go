package v1

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/he-wen-yao/my-blog/server/model/entity"
	"github.com/he-wen-yao/my-blog/server/model/req"
	"github.com/he-wen-yao/my-blog/server/model/res"
	"github.com/he-wen-yao/my-blog/server/service"
)

// ArticleApi 文章相关 api 集合
type ArticleApi struct{}

// CreateArticle 添加文章
func (ArticleApi) CreateArticle(c *gin.Context) {
	var article entity.Article
	err := c.ShouldBindJSON(&article)
	if err != nil {
		res.ResponseGen.FailWithMessage("提交的格式有误", c)
		return
	}
	if len(strings.Trim(article.Title, "")) == 0 {
		res.ResponseGen.FailWithMessage("标题不能为空", c)
		return
	}
	if len(strings.Trim(article.MdText, "")) == 0 {
		res.ResponseGen.FailWithMessage("内容不能为空", c)
		return
	}

	if len(strings.Trim(article.Description, "")) == 0 {
		res.ResponseGen.FailWithMessage("简介不能为空", c)
		return
	}

	if !service.ArticleService.CreateArticle(&article) {
		res.ResponseGen.FailWithMessage("插入失败", c)
		return
	}
	res.ResponseGen.OkWithDetailed(article, "插入成功", c)
}

// UpdateArticle 更新文章
func (ArticleApi) UpdateArticle(c *gin.Context) {
	var article entity.Article
	err := c.ShouldBindJSON(&article)
	if err != nil {
		res.ResponseGen.FailWithMessage("提交的格式有误", c)
		return
	}
	if len(strings.Trim(article.Title, "")) == 0 {
		res.ResponseGen.FailWithMessage("标题不能为空", c)
		return
	}
	if len(strings.Trim(article.MdText, "")) == 0 {
		res.ResponseGen.FailWithMessage("内容不你能为空", c)
		return
	}

	if res_, _ := service.ArticleService.UpdateArticle(&article); !res_ {
		res.ResponseGen.FailWithMessage("更新失败", c)
		return
	}
	res.ResponseGen.OkWithDetailed(article, "更新成功", c)
}

// FetchArticleDetail 获取文章详情
func (ArticleApi) FetchArticleDetail(c *gin.Context) {
	id := c.Param("id")
	if id, err := strconv.ParseInt(id, 10, 64); err == nil {
		article := service.ArticleService.FetchArticleById(id)
		res.ResponseGen.OkWithDetailed(article, "获取成功", c)
		return
	}
	res.ResponseGen.FailWithMessage("文章 ID 填写有误", c)
}

// FetchArticleListByPaging 分页获取文章列表
func (ArticleApi) FetchArticleListByPaging(c *gin.Context) {
	var param req.PagingParam
	err := c.ShouldBindQuery(&param)
	if err != nil {
		res.ResponseGen.FailWithMessage("提交的格式有误", c)
		return
	}
	// 校验参数
	if err := param.Check(); err != nil {
		res.ResponseGen.FailWithMessage(err.Error(), c)
	}
	list, _ := service.ArticleService.ArticleList(&param)
	total := service.ArticleService.FetchTotalNum()
	// 响应
	res.ResponseGen.OkWithData(res.GenPagingResult(list, param, int(total)), c)
}

// DeleteArticleById 删除文章
func (ArticleApi) DeleteArticleById(c *gin.Context) {
	// 获取 ID
	id := c.Param("id")
	if id, err := strconv.ParseInt(id, 10, 64); err == nil {
		res_, _ := service.ArticleService.DeleteArticleById(id)
		if res_ {
			res.ResponseGen.OkWithMessage("删除成功", c)
		} else {
			res.ResponseGen.FailWithMessage("删除失败", c)
		}
		return
	}
	res.ResponseGen.FailWithMessage("文章 ID 填写有误", c)
}
