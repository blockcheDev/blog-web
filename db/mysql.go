package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/express?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Article{})
	DB.AutoMigrate(&ArticleTag{})
	DB.AutoMigrate(&ArticleCategory{})
	DB.AutoMigrate(&Tag{})
	DB.AutoMigrate(&Category{})
	DB.AutoMigrate(&Comment{})
	// db.Create(&User{Name: "blockche", Password: "123456"})
	// DB.Create(&Article{Content: "# test"})
	// DB.Create(&Category{Name: "默认分类"})
}
