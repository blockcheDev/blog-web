package controller

import (
	"net/http"
	"time"
	"webback/db"
	"webback/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ModifyUserInfo(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("token"))
	name := claim.Name
	user := db.User{}
	db.DB.Where("name=?", name).First(&user)
	c.ShouldBind(&user)
	db.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"msg": "修改成功",
	})
}
func ModifyPassword(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("token"))
	name := claim.Name
	form := struct {
		OldPassword string
		NewPassword string
	}{}
	c.ShouldBind(&form)

	// if form.OldPassword != user.Password {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"msg": "原密码错误",
	// 	})
	// 	return
	// }
	user := util.AuthUserAndPassword(name, form.OldPassword)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "原密码错误",
		})
		return
	}

	// 密码加盐后才存入数据库
	hash_password := util.HashPassword(form.NewPassword)

	err := db.DB.Model(&user).Where("name=?", name).Update("password", hash_password).Error
	if err != nil {
		logrus.Error("mysql修改用户密码失败: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "修改密码失败，请稍后再试",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "密码修改成功",
	})
}

func ModifyArticle(c *gin.Context) {
	data := db.Article{}
	// c.ShouldBindBodyWith(&data, binding.JSON) //多次绑定时用这个
	c.ShouldBind(&data)

	claim, _ := util.ParseToken(c.GetHeader("token"))
	name := claim.Name
	article := *db.GetArticle(data.ID)
	if article.UserID != db.GetUserByName(name).ID {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "只能修改本人的文章",
		})
		return
	}

	article.Title = data.Title
	article.Content = data.Content
	article.Type = data.Type
	article.CategoryID = data.CategoryID
	article.ModifiedAt = time.Now()
	db.DB.Save(&article)

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

	c.JSON(http.StatusOK, gin.H{
		"msg": "修改成功",
		"ID":  article.ID,
	})
}
