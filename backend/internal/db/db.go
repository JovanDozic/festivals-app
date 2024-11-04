package db

import (
	modelsCommon "backend/internal/models/common"
	modelsUser "backend/internal/models/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(dbConfig struct{ ConnectionString string }) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(dbConfig.ConnectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// auto migrations go here
	migrateCommonModels(db)
	migrateUserModels(db)

	return db, nil
}

func migrateUserModels(db *gorm.DB) {
	db.AutoMigrate(&modelsUser.User{})
	db.AutoMigrate(&modelsUser.UserProfile{})
	db.AutoMigrate(&modelsUser.Attendee{})
	db.AutoMigrate(&modelsUser.Employee{})
	db.AutoMigrate(&modelsUser.Organizer{})
	db.AutoMigrate(&modelsUser.Log{})
}

func migrateCommonModels(db *gorm.DB) {
	db.AutoMigrate(&modelsCommon.Country{})
	db.AutoMigrate(&modelsCommon.City{})
	db.AutoMigrate(&modelsCommon.Address{})
	db.AutoMigrate(&modelsCommon.Image{})
}
