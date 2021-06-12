package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type envConfig struct {
	SERVER_PORT       string
	DOMAIN            string
	POSTGRES_HOST     string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_PORT     string
	POSTGRES_DB       string
	POSTGRES_SSLMODE  string
	JWT_SECRET        []byte
	JWT_ISSUER        string
}

var config *envConfig

func GetEnv() *envConfig {
	if config == nil {
		config = &envConfig{
			SERVER_PORT:       os.Getenv("SERVER_PORT"),
			DOMAIN:            os.Getenv("DOMAIN"),
			POSTGRES_HOST:     os.Getenv("POSTGRES_HOST"),
			POSTGRES_USER:     os.Getenv("POSTGRES_USER"),
			POSTGRES_PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
			POSTGRES_PORT:     os.Getenv("POSTGRES_PORT"),
			POSTGRES_DB:       os.Getenv("POSTGRES_DB"),
			POSTGRES_SSLMODE:  os.Getenv("POSTGRES_SSLMODE"),
			JWT_SECRET:        []byte(os.Getenv("JWT_SECRET")),
			JWT_ISSUER:        os.Getenv("JWT_ISSUER"),
		}
	}
	return config
}

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func InitOs() {
	initEnv()
	os.Setenv("TZ", "0")
}
