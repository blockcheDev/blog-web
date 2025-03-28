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
