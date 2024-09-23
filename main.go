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
	config.DB.AutoMigrate(&models.Task{}, &models.Comment{})
	routes.TaskRoutes(app)
	routes.CommentRoutes(app)

	app.Listen(":3000")
}
