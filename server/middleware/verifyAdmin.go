package middleware

import (
	"net/http"
	"webback/db"
	"webback/util"

	"github.com/gin-gonic/gin"
)

func VerifyAdmin(c *gin.Context) {
	claim, err := util.ParseToken(c.GetHeader("token"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "token解析失败",
		})
		c.Abort()
	}
	user := db.GetUserByName(claim.Name)

	if !user.VerifyAdmin() {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "权限不足",
		})
		c.Abort()
	}
	c.Next()
}
