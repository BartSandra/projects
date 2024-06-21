package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port      string
	DBSource  string
	JWTSecret string
}

var AppConfig Config

func Load() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	AppConfig.Port = viper.GetString("PORT")
	AppConfig.DBSource = viper.GetString("DB_SOURCE")
	AppConfig.JWTSecret = viper.GetString("JWT_SECRET")
}
