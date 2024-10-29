package proto

import (
	"context"
	"location-service/internal/models"
	pb "location-service/internal/proto/location"
	"location-service/internal/services"
	"log"
)

type Server struct {
	pb.UnimplementedLocationServiceServer
	Service services.LocationService
}

func (p *Server) SaveAddress(ctx context.Context, req *pb.SaveAddressRequest) (*pb.SaveAddressResponse, error) {

	log.Printf("Received address from user: %s, %s, %s, %s, %s, %s", req.Street, req.Number, req.ApartmentSuite, req.City, req.PostalCode, req.Country)

	address := &models.Address{
		Street:         req.Street,
		Number:         req.Number,
		ApartmentSuite: req.ApartmentSuite,
		City: models.City{
			Name:       req.City,
			PostalCode: req.PostalCode,
			Country: models.Country{
				Name:     req.Country,
				ISOCode3: "TEMP",
			},
		},
	}

	if err := p.Service.CreateAddress(address, &address.City, &address.City.Country); err != nil {
		log.Println("error:", err)
		return &pb.SaveAddressResponse{
			Success: false,
		}, err
	}

	log.Println("Address saved successfully")
	return &pb.SaveAddressResponse{
		Success: true,
		Id:      "123",
	}, nil
}
