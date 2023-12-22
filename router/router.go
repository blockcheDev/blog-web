package router

import (
	"webback/controller"
	"webback/middleware"

	"github.com/gin-gonic/gin"
)

func StartRouter() {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.LoadHTMLGlob("./html/*")

	// r.GET("/", controller.ToHome)
	// r.GET("/home", func(c *gin.Context) {
	// 	name, err := c.Cookie("Name")
	// 	if err != nil {
	// 		name = ""
	// 	}
	// 	c.HTML(http.StatusOK, "home.tmpl", gin.H{"name": name})
	// })
	// r.GET("/register", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "register.tmpl", gin.H{})
	// })
	r.POST("/register", controller.Register)
	// r.GET("/login", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "login.tmpl", gin.H{})
	// })
	r.POST("/login", controller.Login)
	// r.GET("/logout", func(c *gin.Context) {
	// 	c.SetCookie("Name", "", -1, "/", "localhost", false, true)
	// 	c.SetCookie("Password", "", -1, "/", "localhost", false, true)
	// 	// c.JSON(http.StatusOK, gin.H{
	// 	// 	"msg": "登出成功",
	// 	// })
	// 	c.Redirect(http.StatusFound, "/home")
	// })

	r.POST("/auth", controller.AuthLoginStatus)

	user := r.Group("/user")
	user.Use(middleware.AuthLogin)
	{
		user.GET("", controller.GetUserInfo)
		user.POST("/update", controller.UpdateUserInfo)
		user.POST("/password", controller.UpdatePassword)
		user.POST("/push", controller.PushArticle)
		user.POST("/create/tag", controller.CreateTag)
		user.POST("/create/category", controller.CreateCategory)
	}

	// article := r.Group("/article")
	// {
	// }

	getInfo := r.Group("/getInfo")
	{
		getInfo.GET("/category/:id", controller.GetCategory)
		getInfo.GET("/tag/:id", controller.GetTag)
		getInfo.GET("/article/:id", controller.GetArticle)
		getInfo.GET("/article/:id/tag", controller.GetTagByArticle)
	}

	// r.GET("/usertest", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"Name":      "blockche",
	// 		"Gender":    "男",
	// 		"Telephone": "10086",
	// 	})
	// })

	// r_user := r.Group("/user")
	// {
	// 	// 中间件：验证登录状态
	// 	r_user.Use(middleware.AuthLogin)
	// 	r_user.GET("/", func(c *gin.Context) {
	// 		u := &db.User{}
	// 		name, _ := c.Cookie("Name")
	// 		db.DB.Where("Name = ?", name).First(u)
	// 		u.Password = "不可视"
	// 		c.JSON(http.StatusOK, u)
	// 	})
	// }

	r.Run(":8080")
}
