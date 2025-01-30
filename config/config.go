package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DNS string
}

func NewAppConfig() *AppConfig {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err.Error())
	}
	return &AppConfig{
		DNS: os.Getenv("DNS"),
	}
}
