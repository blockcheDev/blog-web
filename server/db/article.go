package db

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ModifiedAt time.Time `gorm:"not null"` // 文章修改时间

	CategoryID uint `gorm:"not null; index:category_id"`
	UserID     uint `gorm:"not null"`

	Title     string `gorm:"not null"`
	Content   string `gorm:"not null"`
	Type      int8   `gorm:"not null"` //0-原创 1-转载
	Tags      []uint `gorm:"-"`
	PageViews uint64 `gorm:"default:0"` // 文章浏览量
}
type Articles []Article

type ArticleTag struct {
	gorm.Model
	ArticleID uint `gorm:"not null; index:idx1,priority:1; index:idx2,priority:2"`
	TagID     uint `gorm:"not null; index:idx1,priority:2; index:idx2,priority:1"`
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

func (article *Article) IncreasePageViews() {
	article.PageViews++
	DB.Save(article)
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
