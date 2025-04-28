package db

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type WebVisitor struct {
	gorm.Model
	IP            string    `gorm:"primaryKey"`
	RecentVisitAt time.Time `gorm:"not null; index=recent_visit_at"`
	Count         uint64    `gorm:"default:0; index=count"` // 访问(后端接口)次数
}

func InsertWebVisitor(visitor *WebVisitor) {
	db_visitor := WebVisitor{}
	DB.Where("ip=?", visitor.IP).First(&db_visitor)

	db_visitor.IP = visitor.IP
	db_visitor.RecentVisitAt = visitor.RecentVisitAt
	db_visitor.Count++
	DB.Save(&db_visitor)
}

// func LoadWebVisitorFromRedis() {
// 	list_rdb_visitor, err := RDB.ZRangeWithScores(context.Background(), "recent_visitors", 0, -1).Result()
// 	if err != nil {
// 		logrus.Error("RDB.ZRangeWithScores failed: ", err)
// 		return
// 	}

// 	var db_count int64
// 	err = DB.Model(&WebVisitor{}).Count(&db_count).Error
// 	if err != nil {
// 		logrus.Error("DB.Count failed: ", err)
// 		return
// 	}

// 	if db_count < int64(len(list_rdb_visitor)) {
// 		for _, rdb_visitor := range list_rdb_visitor {
// 			InsertWebVisitor(&WebVisitor{
// 				IP:            rdb_visitor.Member.(string),
// 				RecentVisitAt: time.Unix(int64(rdb_visitor.Score), 0),
// 			})
// 		}
// 	}
// }

func GetWebVisitorCount() int64 {
	var count int64
	err := DB.Model(&WebVisitor{}).Count(&count).Error
	if err != nil {
		logrus.Error("获取网站总访客量失败 ", err)
		return 0
	}
	return count
}
