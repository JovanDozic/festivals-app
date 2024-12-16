package common

import (
	"backend/internal/config"
	dto "backend/internal/dto/common"
	"context"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type AWSService interface {
	GetPresignedURL(req *dto.GetPresignedURLRequest) (*dto.GetPresignedURLResponse, error)
}

type awsService struct {
	s3Client  *s3.Client
	s3Presign *s3.PresignClient
	config    *config.Config // todo: maybe we don't need this as we're injecting clients
}

func NewAWSService(s3Client *s3.Client, s3Presign *s3.PresignClient, config *config.Config) AWSService {
	return &awsService{
		s3Client:  s3Client,
		s3Presign: s3Presign,
		config:    config,
	}
}

func (s *awsService) GetPresignedURL(req *dto.GetPresignedURLRequest) (*dto.GetPresignedURLResponse, error) {

	if !isImage(req.Filename) {
		return nil, fmt.Errorf("file type not supported: %s", req.FileType)
	}

	uniqueFileName := fmt.Sprintf("%s-%s", req.Filename, uuid.New().String())

	objectKey := uniqueFileName

	presignDuration := 15 * time.Minute

	presignParams := &s3.PutObjectInput{
		Bucket:      aws.String(s.config.AWS.S3BucketName),
		Key:         aws.String(objectKey),
		ContentType: aws.String(req.FileType),
	}

	presignedReq, err := s.s3Presign.PresignPutObject(context.Background(), presignParams, func(po *s3.PresignOptions) {
		po.Expires = presignDuration
	})
	if err != nil {
		// todo: remove logging from here
		log.Println("Failed to sign request", err)
		return nil, err
	}

	imageURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.config.AWS.S3BucketName, s.config.AWS.Region, objectKey)

	response := dto.GetPresignedURLResponse{
		UploadURL: presignedReq.URL,
		ImageURL:  imageURL,
	}

	return &response, nil
}

func isImage(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".tiff":
		return true
	default:
		return false
	}
}
