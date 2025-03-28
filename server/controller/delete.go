package controller

import (
	"net/http"
	"webback/db"
	"webback/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func DeleteUser(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("token"))
	name := claim.Name

	user := &db.User{}
	err := c.ShouldBind(user)
	if err != nil {
		logrus.Error("err:", err.Error())
	}

	// res := db.DB.Where("name=? and password=?", name, user.Password).First(&user)
	user = util.AuthUserAndPassword(name, user.Password)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "密码错误",
		})
		return
	}

	db.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{
		"msg": "注销成功",
	})
}

func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	article := *db.GetArticle(id)
	claim, _ := util.ParseToken(c.GetHeader("token"))
	name := claim.Name
	if article.UserID != db.GetUserByName(name).ID {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "只能删除自己的文章",
		})
		return
	}

	//删除关联分类
	db.DB.Where("article_id = ?", id).Delete(&db.ArticleCategory{})
	//删除关联标签
	db.DB.Where("article_id = ?", id).Delete(&db.ArticleTag{})

	db.DB.Delete(&article)
	c.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
	})
}

func DeleteComment(c *gin.Context) {
	id := c.Param("id")
	comment := *db.GetComment(id)
	claim, _ := util.ParseToken(c.GetHeader("token"))
	name := claim.Name
	if comment.UserID != db.GetUserByName(name).ID {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "只能删除自己的评论",
		})
		return
	}

	db.DB.Delete(&comment)
	c.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
	})
}
