package main

import (
	"auth-serivce/internal/config"
	"auth-serivce/internal/db"
	"log"
)

func main() {

	var config config.Config
	config.Load()
	if err := config.Validate(); err != nil {
		log.Fatalln("error loading configuration:", err)
	}
	log.Println("config loaded successfully")

	db, err := db.Init(config.DB)
	if err != nil {
		log.Fatalln("error initializing database:", err)
	}
	log.Println("database initialized successfully")

	// todo: remove this line
	log.Println(db)

	// todo: router

	// todo: start server
}
