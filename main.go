package main

import (
	"todo-list/config"
	"todo-list/models"
	"todo-list/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Task{})
	routes.TaskRoutes(app)

	app.Listen(":3000")
}
