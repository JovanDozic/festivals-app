package proto

import (
	"context"
	pb "location-service/internal/proto/location"
	"log"
)

type Server struct {
	pb.UnimplementedLocationServiceServer
}

func (p *Server) SaveAddress(ctx context.Context, req *pb.SaveAddressRequest) (*pb.SaveAddressResponse, error) {

	log.Printf("Received address from user: %s, %s, %s, %s, %s, %s", req.Street, req.Number, req.ApartmentSuite, req.City, req.PostalCode, req.Country)

	// todo: save address to database

	return &pb.SaveAddressResponse{
		Success: true,
		Id:      "123",
	}, nil
}
