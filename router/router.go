package router

import (
	"net/http"
	"webback/controller"
	"webback/middleware"

	"github.com/gin-gonic/gin"
)

func StartRouter() {
	r := gin.Default()
	r.Use(middleware.Cors())
	// r.LoadHTMLGlob("./html/*")
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "成功访问后端",
		})
	})
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	r.GET("/category/:id", controller.GetArticleListByCategory)
	r.GET("/tag/:id", controller.GetArticleListByTag)
	r.GET("/article/:id/comment", controller.GetCommentListByArticle)

	// r.POST("/auth", controller.AuthLoginStatus)

	user := r.Group("/user")
	user.Use(middleware.AuthLogin)
	{
		user.GET("", controller.GetUserInfo)
		user.POST("/update", controller.UpdateUserInfo)
		user.POST("/password", controller.UpdatePassword)
		user.POST("/push", controller.PushArticle)
		user.POST("/modify/article", controller.ModifyArticle)
		user.POST("/create/tag", controller.CreateTag)
		user.POST("/create/category", controller.CreateCategory)
		user.POST("/create/comment", controller.PushComment)

		user.POST("/delete/user", controller.DeleteUser)
		user.POST("/delete/article/:id", controller.DeleteArticle)
		user.DELETE("/delete/comment/:id", controller.DeleteComment)
	}

	getInfo := r.Group("/getInfo")
	{
		getInfo.GET("/category/:id", controller.GetCategory)
		getInfo.GET("/tag/:id", controller.GetTag)
		getInfo.GET("/article/:id", controller.GetArticle)
		getInfo.GET("/article/:id/tag", controller.GetTagByArticle)
		getInfo.GET("/article/:id/tagid", controller.GetTagIDByArticle)
		getInfo.GET("/user/:id", controller.GetUser)
	}

	r.Run(":8080")
}
