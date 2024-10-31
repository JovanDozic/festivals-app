package proto

import (
	"context"
	pb "location-service/internal/proto/location"
	"location-service/internal/services"
	"log"
)

type Server struct {
	pb.UnimplementedLocationServiceServer
	Service services.LocationService
}

func (p *Server) SaveAddress(ctx context.Context, req *pb.SaveAddressRequest) (*pb.SaveAddressResponse, error) {

	log.Println("Received address from user")

	addressId, err := p.Service.CreateAddress(req)
	if err != nil {
		log.Println("error creating new address:", err)
		return &pb.SaveAddressResponse{
			Success: false,
		}, err
	}

	log.Println("Address saved successfully")
	return &pb.SaveAddressResponse{
		Success: true,
		Id:      addressId.String(),
	}, nil
}
