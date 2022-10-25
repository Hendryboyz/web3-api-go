package configs

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

func ConfigServer(environment string) *viper.Viper {
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(environment)
	config.AddConfigPath("../configs/")
	config.AddConfigPath("configs/")
	if err := config.ReadInConfig(); err != nil {
		log.Fatal("Parse configuration fail")
	}
	return config
}

func GetConfig() *viper.Viper {
	return config
}
