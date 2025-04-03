package main

import (
	"io"
	"os"
	"webback/config"
	"webback/db"
	"webback/router"
	"webback/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func log_init() {
	os.Mkdir("logs", 0755)
	log_file, err := os.OpenFile("logs/today.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatal("Failed to open log file:", err)
	}
	writer := io.MultiWriter(os.Stdout, log_file)
	logrus.SetOutput(writer)
	gin.DefaultWriter = writer

	// 启动日志归档异步任务
	util.LogArchive()
}

func main() {
	// 初始化日志
	log_init()

	// 初始化配置，配置文件在./config/config.yaml，如果没有请参考viper.go文件自行创建
	err := config.Init()
	if err != nil {
		logrus.Error("config.Init failed, err:", err)
		return
	}

	// 连接mysql
	err = db.InitMysql()
	if err != nil {
		logrus.Error("db.InitMysql failed, err:", err)
		return
	}

	// 连接redis
	err = db.InitRedis()
	if err != nil {
		logrus.Error("db.InitRedis failed, err:", err)
		return
	}

	// 启动路由，然后阻塞程序
	router.StartRouter()
}
