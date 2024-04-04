package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBPort     string `mapstructure:"DB_PORT"`
}

var env = []string{"DB_HOST", "DB_NAME", "DB_USER", "DB_PASSWORD", "DB_PORT"}

func LoadConfig() (Config, error) {
	var config Config

	for _, env := range env {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env file loading error due to : ", err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}
	return config, nil
}
