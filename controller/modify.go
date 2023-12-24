package controller

import (
	"net/http"
	"webback/db"
	"webback/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func ModifyArticle(c *gin.Context) {
	data := db.Article{}
	c.ShouldBindBodyWith(&data, binding.JSON) //多次绑定时用这个
	if data.UserID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "禁止修改文章作者",
		})
		return
	}

	claim, _ := util.ParseToken(c.GetHeader("token"))
	name := claim.Name
	article := *db.GetArticle(data.ID)
	if article.UserID != db.GetUserByName(name).ID {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "只能修改本人的文章",
		})
		return
	}

	//修改分类
	db.DB.Where("article_id = ?", article.ID).Delete(&db.ArticleCategory{})
	db.DB.Create(&db.ArticleCategory{
		ArticleID:  data.ID,
		CategoryID: data.CategoryID,
	})
	//修改标签
	db.DB.Where("article_id = ?", article.ID).Delete(&db.ArticleTag{})
	for _, ID := range data.Tags {
		db.DB.Create(&db.ArticleTag{
			ArticleID: data.ID,
			TagID:     ID,
		})
	}
	c.ShouldBindBodyWith(&article, binding.JSON)
	db.DB.Save(&article)
	c.JSON(http.StatusOK, gin.H{
		"msg": "修改成功",
		"ID":  article.ID,
	})
}
