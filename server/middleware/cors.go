package middleware

import (
	"net"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetRealIP(c *gin.Context) string {
	// 优先从 X-Forwarded-For 获取
	xForwardedFor := c.GetHeader("X-Forwarded-For")
	if xForwardedFor != "" {
		ips := strings.Split(xForwardedFor, ",")
		logrus.Info("get from X-Forwarded-For: ", ips)
		return strings.TrimSpace(ips[0]) // 取第一个非代理 IP
	}

	// 次优先从 X-Real-IP 获取
	if realIP := c.GetHeader("X-Real-IP"); realIP != "" {
		logrus.Info("get from X-Real-IP: ", realIP)
		return realIP
	}

	// 兜底方案：从请求连接获取
	ip, _, _ := net.SplitHostPort(c.Request.RemoteAddr)
	return ip
}

// 处理跨域禁止访问的问题
func Cors() gin.HandlerFunc {
	// config := cors.DefaultConfig()
	// // config.AllowOrigins = []string{"http://localhost:5173"} //允许从前端的域访问后端
	// config.AllowAllOrigins = true
	// config.AllowHeaders = []string{"token", "Content-Type", "Authorization"} // 允许的请求头
	// config.AllowCredentials = true
	// return cors.New(config)

	return func(c *gin.Context) {

		logrus.Infof("real remote ip addr: %s", GetRealIP(c))

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
