package config

import (
	"github.com/spf13/viper"
)

func Init() {
	viper.SetDefault("ws.addr", ":8080")
	viper.SetDefault("log.level", 4)

	viper.SetConfigName("config")
	viper.SetConfigName("dev.config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.SetConfigType("json")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()

	viper.SetEnvPrefix("VC")
	viper.AutomaticEnv()
}
