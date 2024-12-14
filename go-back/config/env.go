package config

import (
	"mobarter/app"
	"os"

	"github.com/joho/godotenv"
)

type AppEnv struct {
	Host     string
	Port     string
	Password string
	User     string
	SSLMode  string
	DbName   string
	DbUrl    string
}

var DbParams = AppEnv{
	Host:     os.Getenv("DB_HOST"),
	Port:     os.Getenv("DB_PORT"),
	Password: os.Getenv("DB_PASSWORD"),
	User:     os.Getenv("DB_USER"),
	SSLMode:  os.Getenv("DB_SSLMODE"),
	DbName:   os.Getenv("DB_NAME"),
	DbUrl:    os.Getenv("DB_URL"),
}

func SetupEnv() *AppEnv {
	if err := godotenv.Load(".env"); err != nil {
		app.ErrorPanic(err, "Could not load env")
	}

	config := &AppEnv{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DbName:   os.Getenv("DB_NAME"),
		DbUrl:    os.Getenv("DB_URL"),
	}
	return config
}
