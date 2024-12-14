package app

import (
	"context"
	"mobarter/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"golang.org/x/exp/slog"
	"gorm.io/gorm"
)

type AppState struct {
	DB        *gorm.DB
	logger    *slog.JSONHandler
	Logger    *slog.Logger
	DbQueries *db.Queries
	Ctx       context.Context
}

func (app AppState) NewApp() *fiber.App {
	server := fiber.New(fiber.Config{
		ServerHeader:      "Harmony Server",
		AppName:           "Harmony",
		Prefork:           false,
		CaseSensitive:     false,
		StrictRouting:     false,
		EnablePrintRoutes: false,
		IdleTimeout:       5,
		ReduceMemoryUsage: true,
	})

	server.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3666",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	return server
}
