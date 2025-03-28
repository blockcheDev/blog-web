package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 处理跨域禁止访问的问题
func Cors() gin.HandlerFunc {
	// config := cors.DefaultConfig()
	// // config.AllowOrigins = []string{"http://localhost:5173"} //允许从前端的域访问后端
	// config.AllowAllOrigins = true
	// config.AllowHeaders = []string{"token", "Content-Type", "Authorization"} // 允许的请求头
	// config.AllowCredentials = true
	// return cors.New(config)

	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "*")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusAccepted)
		}
		c.Next()
	}
}
