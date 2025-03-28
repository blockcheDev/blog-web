package main

import (
	"webback/config"
	"webback/db"
	"webback/router"

	"github.com/sirupsen/logrus"
)

func main() {
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
