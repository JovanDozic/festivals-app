package container

import (
	"backend/internal/config"
	"backend/internal/handlers"
	"backend/internal/repos"
	"backend/internal/services"
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"gorm.io/gorm"
)

type Container struct {
	DB        *gorm.DB
	Config    *config.Config
	S3Client  *s3.Client
	S3Presign *s3.PresignClient
	Repos     *repos.Repos
	Services  *services.Services
	Handlers  *handlers.Handlers
}

func NewContainer(db *gorm.DB, cfg *config.Config) (*Container, error) {

	awsCfg, err := awsConfig.LoadDefaultConfig(context.TODO(),
		awsConfig.WithRegion(cfg.AWS.Region),
		awsConfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(cfg.AWS.AccessKeyID, cfg.AWS.SecretAccessKey, ""),
		),
	)
	if err != nil {
		return nil, err
	}

	s3Client := s3.NewFromConfig(awsCfg)
	s3Presign := s3.NewPresignClient(s3Client)

	repos := repos.NewRepositories(db)

	services := services.NewServices(cfg, repos, s3Client, s3Presign)

	handlers := handlers.NewHandlers(services)

	return &Container{
		DB:        db,
		Config:    cfg,
		S3Client:  s3Client,
		S3Presign: s3Presign,
		Repos:     repos,
		Services:  services,
		Handlers:  handlers,
	}, nil
}
