package controller

import (
	"fmt"
	"net/http"
	"webback/db"
	"webback/util"

	"github.com/gin-gonic/gin"
)

func ToHome(c *gin.Context) {
	c.Redirect(http.StatusFound, "/home")
}

func Login(c *gin.Context) {
	user := db.User{}
	if err := c.ShouldBind(&user); err != nil {
		fmt.Println("err:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "提交错误",
		})
		return
	}
	if user.Name == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "信息为空",
		})
		return
	}

	u := util.AuthUserAndPassword(user.Name, user.Password)
	if u != nil {
		tokenString, err := util.GenerateToken(u.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "生成token失败",
			})
			return
		}
		c.Header("token", tokenString)
		c.JSON(http.StatusOK, gin.H{
			"msg": "登录成功",
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "账户或密码错误",
		})
	}
}

func Register(c *gin.Context) {
	u := &db.User{}
	c.ShouldBind(u)

	if u.Name == "" || u.Password == "" || u.Email == "" || u.Gender == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "注册失败，信息不能为空",
		})
		return
	}
	// if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,})\.([a-z]{2,4})$`, u.Email); !m {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"msg": "注册失败，邮箱格式错误",
	// 	})
	// 	return
	// }

	res := db.DB.Where("name = ?", u.Name).First(&db.User{})
	if res.RowsAffected != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "注册失败，用户名已存在",
		})
		return
	}

	res = db.DB.Create(&u)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "注册失败，信息导入错误",
			"err": "数据库无法写入这条数据",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}

// func AuthLoginStatus(c *gin.Context) {
// 	name, err := c.Cookie("Name")
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"msg": "您未登录",
// 		})
// 		return
// 	}
// 	password, err := c.Cookie("Password")
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"msg": "您未登录",
// 		})
// 		return
// 	}

// 	u := util.AuthUserAndPassword(name, password)
// 	if u != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"msg": "成功",
// 		})
// 	} else {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"msg": "身份验证错误，重新登录",
// 		})
// 	}
// }

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

func UpdateUserInfo(c *gin.Context) {
	fmt.Println(c.Request)
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
func UpdatePassword(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("token"))
	name := claim.Name
	user := db.User{}
	db.DB.Where("name=?", name).First(&user)
	form := struct {
		OldPassword string
		NewPassword string
	}{}
	c.ShouldBind(&form)

	fmt.Println(form)

	if form.OldPassword != user.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "原密码错误",
		})
		return
	}

	user.Password = form.NewPassword
	db.DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{
		"msg": "密码修改成功",
	})
}
