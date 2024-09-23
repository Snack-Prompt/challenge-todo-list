package main

import (
	"todo-list/config"
	"todo-list/models"
	"todo-list/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3001",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Task{})
	routes.TaskRoutes(app)

	app.Listen(":3000")
}
