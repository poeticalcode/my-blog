package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/he-wen-yao/my-blog/server/util"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Blog-Authorization")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"msg": "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}
		// parseToken 解析token包含的信息
		claims, err := util.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
	}
}
