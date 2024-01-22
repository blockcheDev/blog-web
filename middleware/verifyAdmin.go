package middleware

import (
	"net/http"
	"webback/db"
	"webback/util"

	"github.com/gin-gonic/gin"
)

func VerifyAdmin(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("token"))
	user := db.GetUserByName(claim.Name)

	if !user.VerifyAdmin() {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "权限不足",
		})
		c.Abort()
	}
	c.Next()
}
