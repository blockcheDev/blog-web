package logic

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ModifiedAt time.Time

	Category Category
	User     User

	Title     string
	Content   string
	Type      int8 //0-原创 1-转载
	Tags      []Tag
	PageViews uint64
	Likes     uint64 // 文章点赞数
}
type Articles []Article

type Category struct {
	ID   uint
	Name string
}

type User struct {
	ID   uint
	Name string
}

type Tag struct {
	ID   uint
	Name string
}

type Comment struct {
	gorm.Model

	User              User
	ArticleID         uint
	ReplyTopCommentID uint
	ReplyCommentID    uint

	Content string
	Favor   int
}
