package controller

import (
	"net/http"
	"webback/db"
	"webback/util"

	"github.com/gin-gonic/gin"
)

func PushArticle(c *gin.Context) {
	article := db.Article{}
	c.ShouldBind(&article)

	if article.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "标题不能为空",
		})
		return
	}
	if article.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "内容不能为空",
		})
		return
	}
	if article.CategoryID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "分类不能为空",
		})
		return
	}

	claim, _ := util.ParseToken(c.GetHeader("token"))
	name := claim.Name
	user := db.GetUserByName(name)
	article.UserID = user.ID

	res := db.DB.Create(&article)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "文章插入数据库失败",
		})
		return
	}

	//处理分类
	rela := db.ArticleCategory{
		ArticleID:  article.ID,
		CategoryID: article.CategoryID,
	}
	db.DB.Create(&rela)

	//处理标签
	for _, ID := range article.Tags {
		rela := db.ArticleTag{
			ArticleID: article.ID,
			TagID:     ID,
		}
		db.DB.Create(&rela)
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "文章发布成功",
		"ID":  article.ID,
	})

}

func CreateTag(c *gin.Context) {
	tag := db.Tag{}
	c.ShouldBind(&tag)

	if tag.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "名字不能为空",
		})
		return
	}

	res := db.DB.Where("name=?", tag.Name).First(&db.Tag{})
	if res.RowsAffected != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "标签已存在",
		})
		return
	}
	db.DB.Create(&tag)
	c.JSON(http.StatusOK, gin.H{
		"msg": "创建成功",
	})
}
func CreateCategory(c *gin.Context) {
	cate := db.Category{}
	c.ShouldBind(&cate)

	if cate.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "名字不能为空",
		})
		return
	}

	res := db.DB.Where("name=?", cate.Name).First(&db.Category{})
	if res.RowsAffected != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "分类已存在",
		})
		return
	}
	db.DB.Create(&cate)
	c.JSON(http.StatusOK, gin.H{
		"msg": "创建成功",
	})
}
