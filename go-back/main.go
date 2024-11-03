package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hello world!")

	router := fiber.New()
	app := fiber.New()

	app.Mount("/api", router)

	router.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "sucess",
			"message": "welcome",
		})
	})

	log.Fatal(app.Listen(":3333"))

}
