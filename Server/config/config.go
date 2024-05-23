package config

import "github.com/spf13/viper"

var Config *viper.Viper

func GetConfig() *viper.Viper {
	return Config
}
