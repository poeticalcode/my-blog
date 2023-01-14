package client

import (
	"github.com/gin-gonic/gin"
	"github.com/he-wen-yao/my-blog/server/db"
	"github.com/he-wen-yao/my-blog/server/model/do"
	"github.com/he-wen-yao/my-blog/server/model/dto"
	"github.com/he-wen-yao/my-blog/server/service"
	"net/http"
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

	if res := service.AppService.ArticleService.CreateArticle(&article); !res {
		dto.ResponseGen.FailWithMessage("插入失败", c)
		return
	}
	dto.ResponseGen.OkWithDetailed(article, "插入成功", c)
}

// FetchArticleDetail 获取文章详情
func (ArticleApi) FetchArticleDetail(c *gin.Context) {
	id := c.Param("id")
	if id, err := strconv.ParseInt(id, 10, 64); err == nil {
		article := service.AppService.ArticleService.FetchArticleById(id)
		dto.ResponseGen.OkWithDetailed(article, "获取成功", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 4000,
		"msg":  "文章 ID 填写有误",
	})
}

// FetchArticleListByPaging 分页获取文章列表
func (ArticleApi) FetchArticleListByPaging(c *gin.Context) {
	page := c.Query("page")
	pageSize := c.Query("pageSize")

	page_, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 4000,
			"msg":  "page 参数填写有误",
		})
	}
	pageSize_, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 4000,
			"msg":  "pageSize 参数填写有误",
		})
	}
	var article do.Article
	db := db.DB()
	db.Find(&article).Limit(int((page_ - 1) * pageSize_)).Offset(int(pageSize_))
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
		"do":   article,
	})
	return
}
