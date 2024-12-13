package db

import (
	"os"
)

var DbParams = DbConfig{
	Host:     os.Getenv("DB_HOST"),
	Port:     os.Getenv("DB_PORT"),
	Password: os.Getenv("DB_PASSWORD"),
	User:     os.Getenv("DB_USER"),
	SSLMode:  os.Getenv("DB_SSLMODE"),
	DbName:   os.Getenv("DB_NAME"),
	DbUrl:    os.Getenv("DB_URL"),
}

type DbConfig struct {
	Host     string
	Port     string
	Password string
	User     string
	DbName   string
	SSLMode  string
	DbUrl    string
}
