package main

import (
	"backend/internal/config"
	"backend/internal/container"
	"backend/internal/db"
	"backend/internal/router"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {

	// ! Load configuration from ENV

	var config config.Config
	config.Load()
	if err := config.Validate(); err != nil {
		log.Fatalln("error loading configuration:", err)
	}
	log.Println("config loaded successfully")

	// ! Initialize database

	db, err := db.Init(config.DB)
	if err != nil {
		log.Fatalln("error initializing database:", err)
	}
	log.Println("database initialized successfully")

	// ! Initialize container (dependencies)

	container, err := container.NewContainer(db, &config)
	if err != nil {
		log.Fatalln("error initializing container:", err)
	}

	// ! Initialize router

	r := router.Init(container)
	log.Println("router initialized successfully")

	// ! Start server

	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders: []string{
			"Content-Type",
			"Authorization",
		},
	})

	handler := corsOptions.Handler(r)

	log.Println("starting server on port:", config.App.Port)
	if err := http.ListenAndServe(":"+config.App.Port, handler); err != nil {
		log.Println("error starting server:", err)
		panic(err)
	}
}
