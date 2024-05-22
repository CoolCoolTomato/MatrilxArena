package config

import "github.com/spf13/viper"

var Config *viper.Viper

func init() {
	Config = viper.New()
	// 配置Config
	Config.SetConfigName("config")
	Config.SetConfigType("json")
	Config.AddConfigPath(".")
	Config.ReadInConfig()
}

func GetConfig() *viper.Viper {
	return Config
}
