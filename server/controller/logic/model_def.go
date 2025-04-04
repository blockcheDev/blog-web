package logic

import "gorm.io/gorm"

type Article struct {
	gorm.Model

	Category Category
	User     User

	Title     string
	Content   string
	Type      int8 //0-原创 1-转载
	Tags      []Tag
	PageViews uint64
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
