package db

import (
	"fmt"
	"webback/config"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMysql() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Conf.Mysql.User, config.Conf.Mysql.Password, config.Conf.Mysql.Address, config.Conf.Mysql.Database)
	logrus.Infof("dsn:%s", dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Errorf("failed to connect database, err: %v", err)
		return
	}

	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Article{})
	DB.AutoMigrate(&ArticleTag{})
	DB.AutoMigrate(&ArticleCategory{})
	DB.AutoMigrate(&Tag{})
	DB.AutoMigrate(&Category{})
	DB.AutoMigrate(&Comment{})
	return
}
