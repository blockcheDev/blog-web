package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Conf = new(Config)

type Config struct {
	Github `mapstructure:"github"`
	Mysql  `mapstructure:"mysql"`
	Redis  `mapstructure:"redis"`
	JWT    `mapstructure:"jwt"`
}

type Github struct {
	ClientID     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
}

type Mysql struct {
	Address  string `mapstructure:"address"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type Redis struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}

type JWT struct {
	Key string `mapstructure:"key"`
}

func Init() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err = viper.ReadInConfig(); err != nil {
		logrus.Errorf("viper.ReadInConfig failed, err:%v", err)
		return
	}
	if err = viper.Unmarshal(Conf); err != nil {
		logrus.Errorf("viper.Unmarshal failed, err:%v", err)
		return
	}
	logrus.Infof("conf:%#v", Conf)
	return
}
