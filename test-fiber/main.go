package main

import (
	"fiber-clean-arch-demo/config"
	"fiber-clean-arch-demo/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	db := config.InitDB()

	app := fiber.New()
	app.Use(logger.New())

	routes.SetupRoutes(app, db)

	log.Println("Server starting on :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}