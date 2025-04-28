package middleware

import (
	"context"
	"time"
	"webback/db"
	"webback/util"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func IncreaseRecentVisitors(c *gin.Context) {
	go func() {
		time_now := time.Now()
		ip := util.GetRealIP(c)

		err := db.RDB.ZAdd(context.Background(), "recent_visitors", redis.Z{
			Member: ip,
			Score:  float64(time_now.Unix()),
		}).Err()
		if err != nil {
			logrus.Error("redis设置最近访客ip失败: ", err)
		}

		db.InsertWebVisitor(&db.WebVisitor{
			IP:            ip,
			RecentVisitAt: time_now,
		})
	}()

	c.Next()
}
