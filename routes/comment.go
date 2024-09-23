package routes

import (
	"todo-list/controllers"

	"github.com/gofiber/fiber/v2"
)

func CommentRoutes(app *fiber.App) {
	app.Post("/comments", controllers.CreateComment)
	app.Get("/comments/task/:taskID", controllers.GetCommentsByTaskID)
	app.Get("/comments/:id", controllers.GetCommentByID)
	app.Put("/comments/:id", controllers.UpdateComment)
	app.Delete("/comments/:id", controllers.DeleteComment)
}
