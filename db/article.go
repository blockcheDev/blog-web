package db

import (
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
	DB.Where("ID=?", ID).First(&article)
	return &article
}
