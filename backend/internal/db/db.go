package db

import (
	common "backend/internal/common/models"
	user "backend/internal/user/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(dbConfig struct{ ConnectionString string }) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(dbConfig.ConnectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// auto migrations go here
	migrateUserModels(db)
	migrateCommonModels(db)

	// todo: refactor inserting default data
	// db.Exec(`INSERT INTO users (user_id, username, password, email, role) VALUES ('566f1ad2-ec32-4ab2-8feb-0f74c484ed5d', 'jovan', 'jovan', 'jovandozic@gmail.com', 'ADMIN') ON CONFLICT (user_id) DO NOTHING`)

	return db, nil
}

func migrateUserModels(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&user.UserProfile{})
	db.AutoMigrate(&user.Attendee{})
	db.AutoMigrate(&user.Employee{})
	db.AutoMigrate(&user.Organizer{})
}

func migrateCommonModels(db *gorm.DB) {
	db.AutoMigrate(&common.Country{})
	db.AutoMigrate(&common.City{})
	db.AutoMigrate(&common.Address{})
	db.AutoMigrate(&common.Image{})
	db.AutoMigrate(&common.Log{})
}
