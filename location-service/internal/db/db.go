package db

import (
	"location-service/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(dbConfig struct{ ConnectionString string }) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(dbConfig.ConnectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Country{})
	db.AutoMigrate(&models.City{})
	db.AutoMigrate(&models.Address{})

	return db, nil
}
