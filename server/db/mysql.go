package db

import (
	"fmt"
	"time"
	"webback/config"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func InitMysql() (err error) {
	newLogger := logger.New(
		logrus.StandardLogger(),
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Conf.Mysql.User, config.Conf.Mysql.Password, config.Conf.Mysql.Address, config.Conf.Mysql.Database)
	logrus.Infof("dsn:%s", dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
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
