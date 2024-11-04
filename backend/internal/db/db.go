package db

import (
	commonModels "backend/internal/models/common"
	userModels "backend/internal/models/user"

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

	// todo: refactor inserting default data
	// db.Exec(`INSERT INTO users (user_id, username, password, email, role) VALUES ('566f1ad2-ec32-4ab2-8feb-0f74c484ed5d', 'jovan', 'jovan', 'jovandozic@gmail.com', 'ADMIN') ON CONFLICT (user_id) DO NOTHING`)

	return db, nil
}

func migrateUserModels(db *gorm.DB) {
	db.AutoMigrate(&userModels.User{})
	db.AutoMigrate(&userModels.UserProfile{})
	db.AutoMigrate(&userModels.Attendee{})
	db.AutoMigrate(&userModels.Employee{})
	db.AutoMigrate(&userModels.Organizer{})
}

func migrateCommonModels(db *gorm.DB) {
	db.AutoMigrate(&commonModels.Country{})
	db.AutoMigrate(&commonModels.City{})
	db.AutoMigrate(&commonModels.Address{})
	db.AutoMigrate(&commonModels.Image{})
	db.AutoMigrate(&commonModels.Log{})
}
