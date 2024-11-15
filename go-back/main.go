package main

// import (
// 	"fmt"
// 	"log"

// 	"github.com/gofiber/fiber/v2"
// )

// func main() {
// 	fmt.Println("Hello world!")

// 	router := fiber.New()
// 	app := fiber.New()

// 	app.Mount("/api", router)

// 	router.Get("/health", func(c *fiber.Ctx) error {
// 		return c.Status(200).JSON(fiber.Map{
// 			"status":  "sucess",
// 			"message": "welcome",
// 		})
// 	})

// 	log.Fatal(app.Listen(":3333"))

// }

import (
	"log"

	"mobarter/app"
	"mobarter/database"
	"mobarter/helper"

	"os"

	"github.com/joho/godotenv"
	"golang.org/x/exp/slog"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		helper.ErrorPanic(err, "Could not load env")
	}

	config := &database.DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DbName:   os.Getenv("DB_NAME"),
	}

	db, err := database.NewConnection(config)

	if err != nil {
		helper.ErrorPanic(err, "Cannot connect to db")
	}

	database.RunMigrations(db)
	// // Slog
	handlerOpt := &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stderr, handlerOpt))

	slog.SetDefault(logger)

	appState := app.AppState{
		DB: db,
	}

	server := appState.NewApp()

	appState.SetupRoutes(server, logger)

	log.Fatal(server.Listen(":" + os.Getenv("PORT")))
	// log.Fatal()

}
