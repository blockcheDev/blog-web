package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"webback/config"
	"webback/db"
	"webback/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 国际化邮箱正则（支持中文等Unicode字符）
var unicodeEmailRegex = regexp.MustCompile(
	`^[\p{L}0-9.!#$%&'*+/=?^_{|}~-]+` + // 本地部分支持Unicode
		`@([a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+` +
		`[\p{L}]{2,}$`, // 顶级域名支持Unicode
)

func ToHome(c *gin.Context) {
	c.Redirect(http.StatusFound, "/home")
}

func LoginWithGithub(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "登录失败，code为空",
		})
		return
	}
	// 获取access_token
	client_id := config.Conf.Github.ClientID
	client_secret := config.Conf.Github.ClientSecret
	url := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", client_id, client_secret, code)
	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Set("Accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logrus.Errorf("获取access_token失败, err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "登录失败，获取access_token失败",
		})
		return
	}
	github := struct {
		AccessToken string `json:"access_token"`
	}{}
	if err := json.NewDecoder(res.Body).Decode(&github); err != nil {
		logrus.Errorf("解析access_token失败, err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "登录失败，解析access_token失败",
		})
		return
	}
	res.Body.Close()
	logrus.Infof("access_token: %s", github.AccessToken)

	// 获取用户信息
	req, _ = http.NewRequest("GET", "https://api.github.com/user", nil)
	req.Header.Set("Authorization", "token "+github.AccessToken)
	req.Header.Set("Accept", "application/json")
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "登录失败，获取用户信息失败",
		})
		return
	}
	githubUser := struct {
		ID        uint   `json:"id"`
		Login     string `json:"login"`
		AvatarUrl string `json:"avatar_url"`
		Email     string `json:"email"`
	}{}
	if err := json.NewDecoder(res.Body).Decode(&githubUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "登录失败，解析用户信息失败",
		})
		return
	}
	res.Body.Close()
	logrus.Infof("githubUser: %#v", githubUser)
	if githubUser.ID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "登录失败，获取用户信息失败",
		})
		return
	}

	// 开始登录
	name := "github_" + githubUser.Login
	user := db.GetUserByName(name)
	if user == nil {
		// 初次登录，自动注册
		user := &db.User{
			Name:      name,
			Password:  "",
			Email:     githubUser.Email,
			AvatarUrl: githubUser.AvatarUrl,
			Gender:    "未知",
		}
		res := db.DB.Create(&user)
		if res.RowsAffected == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "登录失败，自动注册失败",
			})
			return
		}
	} else {
		// 更新信息
		user.Email = githubUser.Email
		user.AvatarUrl = githubUser.AvatarUrl
		db.DB.Save(&user)
	}
	tokenString, err := util.GenerateToken(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "登录失败，生成token失败",
		})
		return
	}
	c.Header("token", tokenString)
	c.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
	})
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
			"msg": "用户名或密码不能为空",
		})
		return
	}

	if u := util.AuthUserAndPassword(user.Name, user.Password); u != nil {
		tokenString, err := util.GenerateToken(user.Name)
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
			"msg": "信息不能为空",
		})
		return
	}

	if valid := unicodeEmailRegex.MatchString(u.Email); !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "邮箱格式错误",
		})
		return
	}

	if strings.HasPrefix(u.Name, "github_") {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "用户名不能以github_开头",
		})
		return
	}

	res := db.DB.Where("name = ?", u.Name).First(&db.User{})
	if res.RowsAffected != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "用户名已存在",
		})
		return
	}

	// 密码加盐
	u.Password = util.HashPassword(u.Password)

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
