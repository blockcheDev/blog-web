package router

import (
	"net/http"
	"webback/controller"
	"webback/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func StartRouter() {
	r := gin.Default()
	r.Use(middleware.IncreaseRecentVisitors)
	r.Use(middleware.Cors())
	// r.LoadHTMLGlob("./html/*")

	// 无需鉴权的接口
	base := r.Group("/api")
	{
		// 测试是否成功连接后端
		base.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "成功访问后端",
			})
		})
		base.POST("/register", controller.Register)           // 用户注册
		base.POST("/login", controller.Login)                 // 用户登录
		base.GET("/login/github", controller.LoginWithGithub) // 使用github登录

		category := base.Group("/category")
		{
			category.GET("/:id", controller.GetCategory)                   // 获取分类信息
			category.GET("/:id/list", controller.GetArticleListByCategory) // 根据分类id获取文章列表
		}
		tag := base.Group("/tag")
		{
			tag.GET("/:id", controller.GetTag)                   // 获取标签信息
			tag.GET("/:id/list", controller.GetArticleListByTag) // 根据标签id获取文章列表
		}
		article := base.Group("/article")
		{
			article.GET("/:id", controller.GetArticle)                      // 获取文章信息
			article.GET("/:id/comment", controller.GetCommentListByArticle) // 根据文章id获取评论列表
			// article.GET("/:id/tag", controller.GetTagByArticle)             // 根据文章id获取标签名称列表
			// article.GET("/:id/tagid", controller.GetTagIDByArticle)         // 根据文章id获取标签id列表
		}
		user := base.Group("/user")
		{
			user.GET("/:id", controller.GetUserName) // 获取用户名称
		}
		base.GET("/recent_visitors", controller.GetRecentVisitors) // 获取最近访客
	}

	// 需要登录的接口
	base.Use(middleware.AuthLogin)
	{
		user := base.Group("/user")
		{
			user.GET("", controller.GetUserInfo)             // 获取用户信息
			user.PUT("", controller.ModifyUserInfo)          // 修改用户信息
			user.PUT("/password", controller.ModifyPassword) // 修改用户密码
			user.DELETE("", controller.DeleteUser)           // 删除用户
		}
		comment := base.Group("/comment")
		{
			comment.POST("", controller.PushComment)         // 发布新评论
			comment.DELETE("/:id", controller.DeleteComment) // 删除评论
		}

	}
	// 以下是需要管理员权限的接口
	base.Use(middleware.VerifyAdmin)
	{
		article := base.Group("/article")
		{
			article.POST("", controller.PushArticle)         // 发布新文章
			article.PUT("", controller.ModifyArticle)        // 修改文章
			article.DELETE("/:id", controller.DeleteArticle) // 删除文章
		}
		tag := base.Group("/tag")
		{
			tag.POST("", controller.CreateTag) // 创新新标签
		}
		category := base.Group("/category")
		{
			category.POST("", controller.CreateCategory) // 创建新分类
		}
	}

	// r.GET("/test", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": "成功访问后端",
	// 	})
	// })
	// r.POST("/register", controller.Register)
	// r.POST("/login", controller.Login)

	// r.GET("/category/:id", controller.GetArticleListByCategory)
	// r.GET("/tag/:id", controller.GetArticleListByTag)
	// r.GET("/article/:id/comment", controller.GetCommentListByArticle)

	// r.POST("/auth", controller.AuthLoginStatus)

	// user := r.Group("/user")
	// user.Use(middleware.AuthLogin)
	// {
	// 	user.GET("", controller.GetUserInfo)
	// 	user.POST("/update", controller.UpdateUserInfo)
	// 	user.POST("/password", controller.UpdatePassword)
	// 	user.POST("/push", controller.PushArticle)
	// 	user.POST("/modify/article", controller.ModifyArticle)
	// 	user.POST("/create/tag", controller.CreateTag)
	// 	user.POST("/create/category", controller.CreateCategory)
	// 	user.POST("/create/comment", controller.PushComment)

	// 	user.POST("/delete/user", controller.DeleteUser)
	// 	user.POST("/delete/article/:id", controller.DeleteArticle)
	// 	user.DELETE("/delete/comment/:id", controller.DeleteComment)
	// }

	// getInfo := r.Group("/getInfo")
	// {
	// 	getInfo.GET("/category/:id", controller.GetCategory)
	// 	getInfo.GET("/tag/:id", controller.GetTag)
	// 	getInfo.GET("/article/:id", controller.GetArticle)
	// 	getInfo.GET("/article/:id/tag", controller.GetTagByArticle)
	// 	getInfo.GET("/article/:id/tagid", controller.GetTagIDByArticle)
	// 	getInfo.GET("/user/:id", controller.GetUser)
	// }

	r.Run(":8080")
}
