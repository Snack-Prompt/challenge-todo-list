package controllers

import (
	"todo-list/config"
	"todo-list/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var validate = validator.New()

func CreateTask(c *fiber.Ctx) error {
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{"erro": "Entrada inválida"})
	}

	if err := validate.Struct(task); err != nil {
		return c.Status(400).JSON(fiber.Map{"erro": "Dados inválidos", "detalhes": err.Error()})
	}

	task.ID = uuid.New().String()
	config.DB.Create(&task)
	return c.Status(201).JSON(task)
}

func GetTasks(c *fiber.Ctx) error {
	status := c.Query("status")
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 0)

	var tasks []models.Task

	if status != "" && status != "pending" && status != "doing" && status != "completed" {
		return c.Status(400).JSON(fiber.Map{
			"erro": "Status inválido. Os status válidos são 'pending', 'doing' ou 'completed'.",
		})
	}

	query := config.DB

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if limit == 0 {
		query.Find(&tasks)
		return c.JSON(tasks)
	}

	offset := (page - 1) * limit
	query.Limit(limit).Offset(offset).Find(&tasks)

	return c.JSON(tasks)
}

func GetTaskByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var task models.Task
	if err := config.DB.First(&task, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"erro": "Tarefa não encontrada"})
	}
	return c.JSON(task)
}

func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var task models.Task

	if err := config.DB.First(&task, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"erro": "Tarefa não encontrada"})
	}

	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{"erro": "Entrada inválida"})
	}

	if err := validate.Struct(task); err != nil {
		return c.Status(400).JSON(fiber.Map{"erro": "Dados inválidos", "detalhes": err.Error()})
	}

	config.DB.Save(&task)
	return c.JSON(task)
}

func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")

	result := config.DB.Where("id = ?", id).Delete(&models.Task{})
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"erro": "Tarefa não encontrada"})
	}

	return c.SendStatus(204)
}
