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
		BaseURL    string
	}
	DB struct {
		ConnectionString string
	}
	JWT struct {
		Secret string
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
	flag.StringVar(&c.App.BaseURL, "baseURL",
		os.Getenv("APP_BASE_URL"),
		"Base URL")
	flag.StringVar(&c.DB.ConnectionString, "dsn",
		os.Getenv("DB_CONNECTION_STRING"),
		"PostgreSQL DSN")
	flag.StringVar(&c.JWT.Secret, "jwt",
		os.Getenv("JWT_SECRET"),
		"JWT secret")

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
	if c.JWT.Secret == "" {
		return errors.New("JWT secret is required")
	}
	return nil
}
