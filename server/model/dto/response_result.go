package dto

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ERROR   = 4000
	SUCCESS = 2000
)

type response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type responseGen struct{}

func (r responseGen) Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, response{
		code,
		data,
		msg,
	})
}

func (r responseGen) Ok(c *gin.Context) {
	r.Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func (r responseGen) OkWithMessage(message string, c *gin.Context) {
	r.Result(SUCCESS, map[string]interface{}{}, message, c)
}

func (r responseGen) OkWithData(data interface{}, c *gin.Context) {
	r.Result(SUCCESS, data, "查询成功", c)
}

func (r responseGen) OkWithDetailed(data interface{}, message string, c *gin.Context) {
	r.Result(SUCCESS, data, message, c)
}

func (r responseGen) Fail(c *gin.Context) {
	r.Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func (r responseGen) FailWithMessage(message string, c *gin.Context) {
	r.Result(ERROR, map[string]interface{}{}, message, c)
}

func (r responseGen) FailWithDetailed(data interface{}, message string, c *gin.Context) {
	r.Result(ERROR, data, message, c)
}

var ResponseGen = responseGen{}
