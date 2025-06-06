package db

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ModifiedAt time.Time // 文章修改时间

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

type ArticleLike struct {
	gorm.Model
	ArticleID uint   `gorm:"not null; index:idx1,priority:1; index:idx2,priority:2"`
	LikeIP    string `gorm:"not null; index:idx1,priority:2; index:idx2,priority:1"` // 点赞IP
}

func GetArticle(ID interface{}) *Article {
	article := Article{}
	DB.Where("id=?", ID).First(&article)
	return &article
}

func (article *Article) IncreasePageViews(ip string) {
	// 查询redis此ip是否存在或是否有效
	is_existing, err := RDB.Exists(context.Background(), fmt.Sprintf("article:read_ip:%v_%s", article.ID, ip)).Result()
	if err != nil {
		logrus.Error("redis查询ip失败:", err)
		return
	}

	_, err = RDB.SetEx(context.Background(), fmt.Sprintf("article:read_ip:%v_%s", article.ID, ip), "", time.Minute).Result()
	if err != nil {
		logrus.Error("redis设置ip失败:", err)
		return
	}

	if is_existing != 0 {
		return
	}

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

func GetTotalPageViews() uint64 {
	var total uint64
	err := DB.Model(&Article{}).Select("SUM(page_views)").Scan(&total).Error
	if err != nil {
		logrus.Error("获取文章总浏览量失败: ", err)
		return 0
	}
	return total
}

func GetArticleLikeCount(article_id uint) uint64 {
	var count int64
	err := DB.Model(&ArticleLike{}).Where("article_id=?", article_id).Count(&count).Error
	if err != nil {
		logrus.Error("获取文章点赞数失败: ", err)
		return 0
	}
	return uint64(count)
}

func AddArticleLikeRecord(article_id uint, like_ip string) (is_exist bool, err error) {
	var count int64
	err = DB.Model(&ArticleLike{}).Where("article_id=? and like_ip=?", article_id, like_ip).Count(&count).Error
	if err != nil {
		logrus.Error("查询文章点赞记录失败: ", err)
		return false, err
	}
	if count > 0 {
		return true, nil
	}

	like := ArticleLike{
		ArticleID: article_id,
		LikeIP:    like_ip,
	}
	err = DB.Create(&like).Error
	if err != nil {
		logrus.Error("添加文章点赞记录失败: ", err)
		return false, err
	}
	return false, nil
}
