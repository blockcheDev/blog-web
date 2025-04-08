package util

import (
	"net"
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
