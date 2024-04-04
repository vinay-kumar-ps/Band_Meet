package db

import (
	"fmt"

	"github.com/Anandhu4456/band-meet/pkg/config"
	"github.com/Anandhu4456/band-meet/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnection(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBPort, cfg.DBName)
	DB, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err := DB.AutoMigrate(&domain.Admin{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain.User{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain.BandProfile{}); err != nil {
		return DB, err
	}
	return DB, err
}
