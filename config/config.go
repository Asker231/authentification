package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DNS string
	SECRET string
}

func NewAppConfig() *AppConfig {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err.Error())
	}
	return &AppConfig{
		DNS: os.Getenv("DNS"),
		SECRET: os.Getenv("SECRET"),
	}
}
