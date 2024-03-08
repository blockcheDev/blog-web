package db

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model

	CategoryID uint `gorm:"not null"`
	UserID     uint `gorm:"not null"`

	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
	Type    int8   `gorm:"not null"` //0-原创 1-转载
	Tags    []uint `gorm:"-"`
}
type Articles []Article

type ArticleTag struct {
	gorm.Model
	ArticleID uint `gorm:"not null"`
	TagID     uint `gorm:"not null"`
}

type ArticleCategory struct {
	gorm.Model
	ArticleID  uint `gorm:"not null"`
	CategoryID uint `gorm:"not null"`
}

func GetArticle(ID interface{}) *Article {
	article := Article{}
	DB.Where("id=?", ID).First(&article)
	return &article
}

// 在创建、更新文章后，需要删除redis中的文章列表缓存
func (a *Article) AfterSave(tx *gorm.DB) (err error) {
	// 删除redis中的文章列表缓存
	logrus.Info("AfterSave触发成功，删除redis中的文章列表缓存")
	_, _ = RDB.Del(context.Background(), "article:all").Result()
	return
}

// 在删除文章后，需要删除redis中的文章列表缓存
func (a *Article) AfterDelete(tx *gorm.DB) (err error) {
	// 删除redis中的文章列表缓存
	logrus.Info("AfterDelete触发成功，删除redis中的文章列表缓存")
	_, _ = RDB.Del(context.Background(), "article:all").Result()
	return
}

// 将文章列表载入redis
func (articles Articles) LoadIntoRedis() (err error) {
	for _, article := range articles {
		j, err := json.Marshal(article)
		if err != nil {
			return err
		}
		_, err = RDB.ZAdd(context.Background(), "article:all", redis.Z{
			Score:  float64(article.CreatedAt.Unix()),
			Member: string(j),
		}).Result()
		if err != nil {
			return err
		}
	}
	return
}
