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
	DB db
}

type db struct {
	Username string
	Name     string
}

// Load loads all environment variables within .env
// If an environment variable is missing, an error will be returned
func Load() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := Config{}

	// DB
	username, err := getEnv("DB_USERNAME")
	name, err := getEnv("DB_NAME")

	cfg.DB = db{
		Username: username,
		Name:     name,
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
