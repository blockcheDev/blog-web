package controller

import (
	"net/http"
	"slices"
	"webback/controller/logic"
	"webback/db"
	"webback/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetUserInfo(c *gin.Context) {
	token := c.Request.Header.Get("token")
	claim, err := util.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "token解析失败",
		})
		return
	}

	logrus.Info("claim.Name: ", claim.Name)
	name := claim.Name
	u := db.GetUserByName(name)
	if u == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "获取用户信息失败",
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
		if user == nil {
			user = &db.User{}
		}
		c.JSON(http.StatusOK, gin.H{
			"Name": user.Name,
		})
	}
}

func GetArticle(c *gin.Context) {
	id := c.Param("id")
	// id为all时返回所有文章
	if id == "all" {
		articles := logic.Articles{}
		// 先从redis中获取
		err := articles.GetFromRedis()
		if true || err != nil || len(articles) == 0 {
			// 有错误 或者 redis中没有，从mysql中载入缓存
			logrus.Info("有错误 或者 redis中没有，从mysql中载入缓存")
			db_articles := db.Articles{}
			db.DB.Find(&db_articles)

			articles = make(logic.Articles, len(db_articles))
			for i, db_article := range db_articles {
				articles[i] = logic.GetArticle(&db_article)
			}

			err := articles.LoadIntoRedis()
			if err != nil {
				logrus.Error("redis缓存失败, err: ", err)
				// 这里不return
			}

			// 文章列表时间倒序
			slices.Reverse(articles)
		}
		c.JSON(http.StatusOK, articles)
	} else {
		db_article := db.GetArticle(id)
		if db_article == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "文章不存在",
			})
			return
		}
		article := logic.GetArticle(db_article)
		c.JSON(http.StatusOK, article)

		// 增加浏览量
		db_article.IncreasePageViews(util.GetRealIP(c))
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

// func GetTagByArticle(c *gin.Context) {
// 	id := c.Param("id")
// 	data := []string{}
// 	db.DB.Model(&db.ArticleTag{}).Select("tags.name").Joins("join tags on tags.id = article_tags.tag_id").Where("article_tags.article_id=?", id).Find(&data)
// 	c.JSON(http.StatusOK, data)
// }
// func GetTagIDByArticle(c *gin.Context) {
// 	id := c.Param("id")
// 	data := []uint{}
// 	db.DB.Model(&db.ArticleTag{}).Select("tags.id").Joins("join tags on tags.id = article_tags.tag_id").Where("article_tags.article_id=?", id).Find(&data)
// 	c.JSON(http.StatusOK, data)
// }

func GetArticleListByCategory(c *gin.Context) {
	id := c.Param("id")
	data := []db.Article{}
	db.DB.Table("articles").Where("category_id=?", id).Find(&data)

	articles := make(logic.Articles, len(data))
	for i, db_article := range data {
		articles[i] = logic.GetArticle(&db_article)
	}

	c.JSON(http.StatusOK, articles)
}

func GetArticleListByTag(c *gin.Context) {
	id := c.Param("id")
	data := []db.Article{}
	db.DB.Debug().Model(&db.ArticleTag{}).Select("articles.*").Joins("join articles on articles.id = article_tags.article_id").Where("article_tags.tag_id=?", id).Find(&data)

	articles := make(logic.Articles, len(data))
	for i, db_article := range data {
		articles[i] = logic.GetArticle(&db_article)
	}

	c.JSON(http.StatusOK, articles)
}

func GetCommentListByArticle(c *gin.Context) {
	id := c.Param("id")
	data := []db.Comment{}
	db.DB.Where("article_id=?", id).Find(&data)

	comments := make([]logic.Comment, len(data))
	for i, db_comment := range data {
		comments[i] = logic.GetComment(&db_comment)
	}

	c.JSON(http.StatusOK, comments)
}
