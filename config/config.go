package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	Base  Basecfg
	Mongo DBcfg
)

func InitConfig() {
	v := viper.New()
	v.SetConfigFile("./config.yml")
	if err := v.ReadInConfig(); err != nil {
		logrus.Fatal("failed to read in config & ", err.Error())
	}
	cfg := Config{}
	if err := v.Unmarshal(&cfg); err != nil {
		logrus.Fatal("failed to unmarshal config & ", err.Error())
	}
	Base = cfg.Base
	Mongo = cfg.Mongo
}
