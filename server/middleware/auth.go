package middleware

import (
	"net/http"
	"webback/util"

	"github.com/gin-gonic/gin"
)

func AuthLogin(c *gin.Context) {
	path := c.Request.URL.Path
	if path == "/login" || path == "/register" {
		c.Next()
	} else {
		tokenString := c.Request.Header.Get("token")

		_, err := util.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "请登录",
			})
			c.Abort()
		}
		c.Next()
	}
}
