package config

import (
	"fmt"
	"mobarter/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDb(config Config) *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DbHost, config.DbPort, config.DbUsername, config.DbPassword, config.DbName)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})

	helper.ErrorPanic(err)

	fmt.Println("Connected to database successfully")
	return db

}
