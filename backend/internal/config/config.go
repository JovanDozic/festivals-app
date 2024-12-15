package config

import (
	"errors"
	"flag"
	"os"
)

type Config struct {
	App struct {
		APIVersion string
		Name       string
		Port       string
		Env        string
	}
	DB struct {
		ConnectionString  string
		RootAdminPassword string
	}
	JWT struct {
		Secret string
	}
	AWS struct {
		AccessKeyID     string
		SecretAccessKey string
		Region          string
		S3BucketName    string
	}
	SMTP struct {
		Host     string
		Port     int
		Username string
		Password string
		From     string
	}
}

func (c *Config) Load() {

	flag.StringVar(&c.App.APIVersion, "apiVersion",
		os.Getenv("APP_API_VERSION"),
		"API version")
	flag.StringVar(&c.App.Name, "name",
		os.Getenv("APP_NAME"),
		"API server name")
	flag.StringVar(&c.App.Port, "port",
		os.Getenv("APP_PORT"),
		"API server port")
	flag.StringVar(&c.App.Env, "env",
		os.Getenv("APP_ENV"),
		"Environment (dev|stage|prod)")
	flag.StringVar(&c.DB.ConnectionString, "dsn",
		os.Getenv("DB_CONNECTION_STRING"),
		"PostgreSQL DSN")
	flag.StringVar(&c.DB.RootAdminPassword, "rootAdminPassword",
		os.Getenv("ROOT_ADMIN_PASSWORD"),
		"Root admin password")
	flag.StringVar(&c.JWT.Secret, "jwt",
		os.Getenv("JWT_SECRET"),
		"JWT secret")
	flag.StringVar(&c.AWS.AccessKeyID, "awsAccessKeyID",
		os.Getenv("AWS_ACCESS_KEY_ID"),
		"AWS access key ID")
	flag.StringVar(&c.AWS.SecretAccessKey, "awsSecretAccessKey",
		os.Getenv("AWS_SECRET_ACCESS_KEY"),
		"AWS secret access key")
	flag.StringVar(&c.AWS.Region, "awsRegion",
		os.Getenv("AWS_REGION"),
		"AWS region")
	flag.StringVar(&c.AWS.S3BucketName, "awsS3BucketName",
		os.Getenv("AWS_S3_BUCKET_NAME"),
		"AWS S3 bucket name")
	flag.StringVar(&c.SMTP.Host, "smtpHost",
		os.Getenv("SMTP_HOST"),
		"SMTP host")
	flag.IntVar(&c.SMTP.Port, "smtpPort",
		587,
		"SMTP port")
	flag.StringVar(&c.SMTP.Username, "smtpUsername",
		os.Getenv("SMTP_USERNAME"),
		"SMTP username")
	flag.StringVar(&c.SMTP.Password, "smtpPassword",
		os.Getenv("SMTP_PASSWORD"),
		"SMTP password")
	flag.StringVar(&c.SMTP.From, "smtpFrom",
		os.Getenv("SMTP_FROM"),
		"SMTP from")

	flag.Parse()
}

func (c Config) Validate() error {
	if c.App.APIVersion == "" {
		return errors.New("API version is required")
	}
	if c.App.Name == "" {
		return errors.New("API server name is required")
	}
	if c.App.Port == "" {
		return errors.New("API server port is required")
	}
	if c.App.Env == "" {
		return errors.New("environment is required")
	}
	if c.DB.ConnectionString == "" {
		return errors.New("PostgreSQL DSN is required")
	}
	if c.DB.RootAdminPassword == "" {
		return errors.New("root admin password is required")
	}
	if c.JWT.Secret == "" {
		return errors.New("JWT secret is required")
	}
	if c.AWS.AccessKeyID == "" {
		return errors.New("AWS access key ID is required")
	}
	if c.AWS.SecretAccessKey == "" {
		return errors.New("AWS secret access key is required")
	}
	if c.AWS.Region == "" {
		return errors.New("AWS region is required")
	}
	if c.AWS.S3BucketName == "" {
		return errors.New("AWS S3 bucket name is required")
	}
	if c.SMTP.Host == "" {
		return errors.New("SMTP host is required")
	}
	if c.SMTP.Username == "" {
		return errors.New("SMTP username is required")
	}
	if c.SMTP.Password == "" {
		return errors.New("SMTP password is required")
	}
	if c.SMTP.From == "" {
		return errors.New("SMTP from is required")
	}

	return nil
}
