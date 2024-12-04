package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"txnbi-backend/conf"
	"txnbi-backend/pkg/jwt"
)

func AuthUserToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 先在query查询token
		token, is := c.GetQuery("token")
		if !is {
			// 如果没找到，则在postForm中找
			token, is = c.GetPostForm("token")
			if !is {
				c.JSON(http.StatusUnauthorized, gin.H{"statusCode": 1, "message": "未登陆！"})
				c.Abort()
				return
			}
		}
		id, userAccount, err := jwt.ParseUserToken(token, conf.JWTCfg.SignKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"statusCode": 1, "message": err.Error()})
			c.Abort()
			return
		}

		c.Set("userID", id)
		c.Set("userAccount", userAccount)
		c.Next()
	}
}
