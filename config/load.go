package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config hold application config environment variables pulled from env via Load()
// Config should not be created outside of func Load
type Config struct {
	DB             db
	Authentication auth
}

type db struct {
	Username string
	Password string
	Name     string
}

type auth struct {
	Username string
	Password string
}

// Load loads all environment variables within .env
// If an environment variable is missing, an error will be returned
func Load(envPath string) (Config, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := Config{}

	// DB
	username, err := getEnv("DB_USERNAME")
	password, err := getEnv("DB_PASSWORD")
	name, err := getEnv("DB_NAME")

	cfg.DB = db{
		Username: username,
		Name:     name,
		Password: password,
	}

	// Authentication
	au, err := getEnv("AUTH_USERNAME")
	ap, err := getEnv("AUTH_PASSWORD")

	cfg.Authentication = auth{
		Username: au,
		Password: ap,
	}

	return cfg, err
}

func getEnv(name string) (string, error) {
	s := os.Getenv(name)
	if s == "" {
		return s, fmt.Errorf("error parsing ENV var: %s not set", name)
	}

	return s, nil
}
