package controller

import "gorm.io/gorm"

type Article struct {
	gorm.Model

	Category string
	UserID   uint

	Title   string
	Content string
	Type    int8 //0-原创 1-转载
	Tags    []string
}
type Articles []Article
