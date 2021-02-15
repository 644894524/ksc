package common

import (
	"github.com/spf13/viper"
)

const (
	WEB_CONFILE		 = "../config/conf.yaml"
	WEB_CONFTYPE	 = "yaml"
)

func InitViper(){
	viper.SetConfigFile(WEB_CONFILE)
	viper.SetConfigType(WEB_CONFTYPE)
	err := viper.ReadInConfig()
	if err != nil {
		panic("读取配置文件失败")
	}

	// 监控配置文件变化
	viper.WatchConfig()
}

