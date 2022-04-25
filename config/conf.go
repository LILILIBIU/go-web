package config

import (
	"github.com/spf13/viper"
	"log"
)

func InitConfig() {
	viper.SetConfigName("application")
	viper.SetConfigType("toml")
	viper.AddConfigPath("E:\\桌面\\GOspace\\webDemo\\web")
	//viper.SetDefault("redis.port", 6381)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}
}
