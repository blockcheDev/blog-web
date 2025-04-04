package db

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name string `gorm:"not null"`
}

func GetTag(ID interface{}) *Tag {
	tag := Tag{}
	DB.Where("id=?", ID).First(&tag)
	return &tag
}

func GetTagsByArticleID(ID uint) []Tag {
	tags := []Tag{}
	DB.Model(&ArticleTag{}).Select("tags.*").Joins("join tags on tags.id = article_tags.tag_id").Where("article_tags.article_id=?", ID).Find(&tags)
	return tags
}
