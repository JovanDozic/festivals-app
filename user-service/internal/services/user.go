package services

import (
	"context"
	"log"
	"strings"
	"time"
	"user-service/internal/config"
	"user-service/internal/models"
	"user-service/internal/repos"
	"user-service/internal/utils"

	"google.golang.org/grpc"

	pb "user-service/internal/proto/location"
)

type UserService interface {
	Create(user *models.User) error
	Login(username string, password string) (string, error)
	// todo: other methods
	SaveAddress(street, number, apartmentSuite, city, zipCode, country string) error
}

type userService struct {
	repo               repos.UserRepo
	config             *config.Config
	grpcLocationClient pb.LocationServiceClient
}

func NewUserService(r repos.UserRepo, c *config.Config) UserService {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	client := pb.NewLocationServiceClient(conn)

	return &userService{repo: r, config: c, grpcLocationClient: client}
}

func (s *userService) Create(user *models.User) error {

	if err := user.Validate(); err != nil {
		return err
	}

	passwordHash, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = passwordHash

	if err := s.repo.Create(user); err != nil {
		switch {
		case strings.Contains(err.Error(), "duplicate key value"):
			return models.ErrDuplicateUser
		case strings.Contains(err.Error(), "foreign key constraint"):
			return models.ErrRoleNotFound
		default:
			return err
		}
	}

	return nil
}

func (s *userService) Login(username string, password string) (string, error) {

	user, err := s.repo.GetByUsername(username)
	if err != nil {
		return "", models.ErrNotFound
	}

	if !utils.VerifyPassword(user.Password, password) {
		return "", models.ErrInvalidPassword
	}

	jwt := utils.NewJWTUtil(s.config.JWT.Secret)
	token, err := jwt.GenerateToken(user.Username, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) SaveAddress(street, number, apartmentSuite, city, postalCode, country string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.SaveAddressRequest{
		Street:         street,
		Number:         number,
		ApartmentSuite: apartmentSuite,
		City:           city,
		PostalCode:     postalCode,
		Country:        country,
	}

	resp, err := s.grpcLocationClient.SaveAddress(ctx, req)
	if err != nil {
		log.Println("error calling SaveAddress:", err)
		return err
	}

	log.Println("response from SaveAddress:", resp)
	return nil
}
