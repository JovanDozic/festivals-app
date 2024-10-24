package main

import (
	"auth-service/internal/config"
	"auth-service/internal/db"
	"auth-service/internal/router"
	"log"
	"net/http"
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
	r := router.Init(db, &config)
	log.Println("router initialized successfully")

	// todo: start server
	log.Println("starting server on port:", config.App.Port, "with base URL:", config.App.BaseURL)
	if err := http.ListenAndServe(":"+config.App.Port, r); err != nil {
		log.Println("error starting server:", err)
		panic(err)
	}
}
