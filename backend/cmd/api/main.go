package main

import (
	"backend/internal/config"
	"backend/internal/db"
	"backend/internal/router"
	"log"
	"net/http"

	"github.com/rs/cors"
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

	r := router.Init(db, &config)
	log.Println("router initialized successfully")

	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Allow all origins, or specify your allowed origins
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
			// Add other headers if needed
		},
	})

	handler := corsOptions.Handler(r)

	log.Println("starting server on port:", config.App.Port)
	if err := http.ListenAndServe(":"+config.App.Port, handler); err != nil {
		log.Println("error starting server:", err)
		panic(err)
	}
}
