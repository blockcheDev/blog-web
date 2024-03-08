package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"webback/db"
	"webback/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetUserInfo(c *gin.Context) {
	claim, err := util.ParseToken(c.Request.Header.Get("token"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "token解析失败",
		})
		return
	}

	name := claim.Name
	u := db.User{}
	res := db.DB.Where("name = ?", name).First(&u)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "获取数据库信息失败",
		})
		return
	}

	c.JSON(http.StatusOK, u)
}
func GetUserName(c *gin.Context) {
	id := c.Param("id")
	if id == "0" {
		claim, err := util.ParseToken(c.Request.Header.Get("token"))
		name := ""
		if err == nil {
			name = claim.Name
		}
		c.JSON(http.StatusOK, gin.H{
			"Name": name,
		})
	} else {
		user := db.GetUser(id)
		c.JSON(http.StatusOK, gin.H{
			"Name": user.Name,
		})
	}
}
func GetArticle(c *gin.Context) {
	id := c.Param("id")
	// id为all时返回所有文章
	if id == "all" {
		// 先从redis中获取
		// _, _ = db.RDB.Del(context.Background(), "article:all").Result()
		res, _ := db.RDB.ZRevRange(context.Background(), "article:all", 0, -1).Result()
		if len(res) == 0 {
			// redis中没有，从mysql中载入缓存
			logrus.Info("redis中没有，从mysql中载入缓存")
			articles := db.Articles{}
			db.DB.Find(&articles)
			articles.LoadIntoRedis()

			// 再从redis中获取一次
			res, _ = db.RDB.ZRevRange(context.Background(), "article:all", 0, -1).Result()
			if len(res) == 0 {
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "获取文章列表失败",
				})
				return
			}
		}
		// 从redis中获取成功
		articles := []db.Article{}
		for _, v := range res {
			article := db.Article{}
			json.Unmarshal([]byte(v), &article)
			articles = append(articles, article)
		}
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
	db.DB.Model(&db.ArticleTag{}).Select("tags.name").Joins("join tags on tags.id = article_tags.tag_id").Where("article_tags.article_id=?", id).Find(&data)
	c.JSON(http.StatusOK, data)
}
func GetTagIDByArticle(c *gin.Context) {
	id := c.Param("id")
	data := []uint{}
	db.DB.Model(&db.ArticleTag{}).Select("tags.id").Joins("join tags on tags.id = article_tags.tag_id").Where("article_tags.article_id=?", id).Find(&data)
	c.JSON(http.StatusOK, data)
}

func GetArticleListByCategory(c *gin.Context) {
	id := c.Param("id")
	data := []db.Article{}
	db.DB.Table("articles").Where("category_id=?", id).Find(&data)
	c.JSON(http.StatusOK, data)
}

func GetArticleListByTag(c *gin.Context) {
	id := c.Param("id")
	data := []db.Article{}
	db.DB.Debug().Model(&db.ArticleTag{}).Select("articles.*").Joins("join articles on articles.id = article_tags.article_id").Where("article_tags.tag_id=?", id).Find(&data)
	c.JSON(http.StatusOK, data)
}

func GetCommentListByArticle(c *gin.Context) {
	id := c.Param("id")
	data := []db.Comment{}
	db.DB.Where("article_id=?", id).Find(&data)
	c.JSON(http.StatusOK, data)
}
