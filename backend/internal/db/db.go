package db

import (
	modelsCommon "backend/internal/models/common"
	modelsFestival "backend/internal/models/festival"
	modelsUser "backend/internal/models/user"
	"backend/internal/utils"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(dbConfig struct{ ConnectionString, RootAdminPassword string }) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(dbConfig.ConnectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	migrateCommonModels(db)
	migrateUserModels(db)
	migrateFestivalModels(db)

	if isTableEmpty(db, "countries") {
		err := runSQLScript(db, "countries.sql")
		if err != nil {
			log.Println("error running countries.sql:", err)
			return nil, err
		}
	}

	insertAdministrator(db, dbConfig.RootAdminPassword)

	return db, nil
}

func migrateUserModels(db *gorm.DB) {
	db.AutoMigrate(&modelsUser.User{})
	db.AutoMigrate(&modelsUser.UserProfile{})
	db.AutoMigrate(&modelsUser.Attendee{})
	db.AutoMigrate(&modelsUser.Employee{})
	db.AutoMigrate(&modelsUser.Organizer{})
	db.AutoMigrate(&modelsUser.Administrator{})
	db.AutoMigrate(&modelsUser.Log{})
}

func migrateCommonModels(db *gorm.DB) {
	db.AutoMigrate(&modelsCommon.Country{})
	db.AutoMigrate(&modelsCommon.City{})
	db.AutoMigrate(&modelsCommon.Address{})
	db.AutoMigrate(&modelsCommon.Image{})
}

func migrateFestivalModels(db *gorm.DB) {
	db.AutoMigrate(&modelsFestival.Festival{})

	db.AutoMigrate(&modelsFestival.FestivalOrganizer{})
	db.AutoMigrate(&modelsFestival.FestivalEmployee{})
	db.AutoMigrate(&modelsFestival.FestivalImage{})

	db.AutoMigrate(&modelsFestival.PriceList{})
	db.AutoMigrate(&modelsFestival.Item{})
	db.AutoMigrate(&modelsFestival.PriceListItem{})

	db.AutoMigrate(&modelsFestival.TicketType{})

	db.AutoMigrate(&modelsFestival.PackageAddon{})
	db.AutoMigrate(&modelsFestival.PackageAddonImage{})

	db.AutoMigrate(&modelsFestival.CampAddon{})
	db.AutoMigrate(&modelsFestival.CampEquipment{})

	db.AutoMigrate(&modelsFestival.TransportAddon{})

	db.AutoMigrate(&modelsFestival.FestivalTicket{})
	db.AutoMigrate(&modelsFestival.FestivalPackage{})
	db.AutoMigrate(&modelsFestival.FestivalPackageAddon{})
	db.AutoMigrate(&modelsFestival.Order{})

	db.AutoMigrate(&modelsFestival.Bracelet{})
	db.AutoMigrate(&modelsFestival.ActivationHelpRequest{})
}

func isTableEmpty(db *gorm.DB, tableName string) bool {
	var count int64
	db.Table(tableName).Count(&count)
	return count == 0
}

func runSQLScript(db *gorm.DB, filePath string) error {
	sqlBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Println("error running SQL script:", err)
		return err
	}

	if err := db.Exec(string(sqlBytes)).Error; err != nil {
		log.Println("failed to execute SQL script: %w", err)
		return err
	}

	return nil
}

func insertAdministrator(db *gorm.DB, password string) error {

	exists := db.Where("username = ?", "admin").First(&modelsUser.User{}).Error
	if exists == nil {
		return nil
	}

	hashedPassword, _ := utils.HashPassword(password)

	user := modelsUser.User{
		Username: "admin",
		Password: hashedPassword,
		Email:    "admin@mock.com",
		Role:     string(modelsUser.RoleAdmin),
	}

	err := db.Create(&user).Error
	if err != nil {
		log.Println("error creating admin account:", err)
	}

	profile := modelsUser.UserProfile{
		FirstName:   "Admin",
		LastName:    "Admin",
		DateOfBirth: time.Now(),
		PhoneNumber: "",
		UserID:      user.ID,
		User:        user,
	}

	err = db.Create(&profile).Error
	if err != nil {
		log.Println("error creating admin:", err)
	}

	admin := modelsUser.Administrator{
		UserID: user.ID,
		User:   user,
	}

	err = db.Create(&admin).Error
	if err != nil {
		log.Println("error creating admin:", err)
	}

	return nil
}
