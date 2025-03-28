package db

import "gorm.io/gorm"

type Comment struct {
	gorm.Model

	UserID            uint `gorm:"not null"`
	ArticleID         uint `gorm:"not null"`
	ReplyTopCommentID uint
	ReplyCommentID    uint

	Content string `gorm:"not null"`
	Favor   int
}

func GetComment(ID interface{}) *Comment {
	comment := Comment{}
	DB.Where("id=?", ID).First(&comment)
	return &comment
}
