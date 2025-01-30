package db

import (
	"github.com/Asker231/authentification.git/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDataBase(config *config.AppConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DNS), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
