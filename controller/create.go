package controller

import (
	"net/http"
	"webback/db"

	"github.com/gin-gonic/gin"
)

func CreateTag(c *gin.Context) {
	tag := db.Tag{}
	c.ShouldBind(&tag)

	if tag.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "名字不能为空",
		})
		return
	}

	res := db.DB.Where("Name=?", tag.Name).First(&db.Tag{})
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

	res := db.DB.Where("Name=?", cate.Name).First(&db.Category{})
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
