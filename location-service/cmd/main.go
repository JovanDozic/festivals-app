package main

import (
	"location-service/internal/config"
	"location-service/internal/db"
	"location-service/internal/proto"
	"log"
	"net"

	pb "location-service/internal/proto/location"

	"google.golang.org/grpc"
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

	lis, err := net.Listen("tcp", ":"+config.App.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterLocationServiceServer(s, &proto.Server{})

	log.Println("starting server on port:", config.App.Port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// // todo: router
	// r := router.Init(db, &config)
	// log.Println("router initialized successfully")

	// // todo: start server
	// log.Println("starting server on port:", config.App.Port, "with base URL:", config.App.BaseURL)
	// if err := http.ListenAndServe(":"+config.App.Port, r); err != nil {
	// 	log.Println("error starting server:", err)
	// 	panic(err)
	// }
}
