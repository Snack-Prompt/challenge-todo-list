package routes

import (
	"todo-list/controllers"

	"github.com/gofiber/fiber/v2"
)

func TaskRoutes(app *fiber.App) {
	app.Post("/tasks", controllers.CreateTask)
	app.Get("/tasks", controllers.GetTasks)
	app.Get("/tasks/:id", controllers.GetTaskByID)
	app.Put("/tasks/:id", controllers.UpdateTask)
	app.Delete("/tasks/:id", controllers.DeleteTask)
}
