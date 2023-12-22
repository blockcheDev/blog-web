package db

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"not null"`
}

func GetCategory(ID interface{}) *Category {
	category := Category{}
	DB.Where("ID=?", ID).First(&category)
	return &category
}
