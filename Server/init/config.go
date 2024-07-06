package init

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/config"
	"github.com/spf13/viper"
)

func ConfigInit() {
	config.Config = viper.New()
	config.Config.SetConfigName("config")
	config.Config.SetConfigType("json")
	config.Config.AddConfigPath(".")
	config.Config.ReadInConfig()
}
