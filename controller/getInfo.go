package controller

import (
	"net/http"
	"webback/db"

	"github.com/gin-gonic/gin"
)

func GetArticle(c *gin.Context) {
	id := c.Param("id")
	if id == "all" {
		articles := []db.Article{}
		db.DB.Find(&articles)
		c.JSON(http.StatusOK, articles)
	} else {
		article := db.GetArticle(id)
		if article == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "文章不存在",
			})
			return
		}
		c.JSON(http.StatusOK, article)
	}
}
func GetCategory(c *gin.Context) {
	id := c.Param("id")
	if id == "all" {
		cate := []db.Category{}
		db.DB.Find(&cate)
		c.JSON(http.StatusOK, cate)
	} else {
		cate := db.GetCategory(id)
		if cate == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "分类不存在",
			})
			return
		}
		c.JSON(http.StatusOK, cate)
	}
}
func GetTag(c *gin.Context) {
	id := c.Param("id")
	if id == "all" {
		tags := []db.Tag{}
		db.DB.Find(&tags)
		c.JSON(http.StatusOK, tags)
	} else {
		tag := db.GetTag(id)
		if tag == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "分类不存在",
			})
			return
		}
		c.JSON(http.StatusOK, tag)
	}
}
func GetTagByArticle(c *gin.Context) {
	id := c.Param("id")
	data := []string{}
	db.DB.Model(&db.ArticleTag{}).Select("tags.name").Joins("join tags on tags.id = article_tags.tag_id").Where("article_tags.article_id=?", id).Scan(&data)
	c.JSON(http.StatusOK, data)
}