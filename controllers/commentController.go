package controllers

import (
	"todo-list/config"
	"todo-list/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateComment(c *fiber.Ctx) error {
	var comment models.Comment
	if err := c.BodyParser(&comment); err != nil {
		return c.Status(400).JSON(fiber.Map{"erro": "Entrada inválida"})
	}

	var task models.Task
	if err := config.DB.First(&task, "id = ?", comment.TaskID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"erro": "Tarefa não encontrada"})
	}

	comment.ID = uuid.New().String()
	if err := config.DB.Create(&comment).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"erro": "Erro ao criar comentário"})
	}

	return c.Status(201).JSON(fiber.Map{
		"id":        comment.ID,
		"content":   comment.Content,
		"taskId":    comment.TaskID,
		"createdAt": comment.CreatedAt,
		"updatedAt": comment.UpdatedAt,
	})
}

func GetCommentsByTaskID(c *fiber.Ctx) error {
	taskID := c.Params("taskId")
	var comments []models.Comment

	if err := config.DB.Preload("Task").Where("task_id = ?", taskID).Find(&comments).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"erro": "Nenhum comentário encontrado para essa tarefa"})
	}

	return c.JSON(comments)
}
func GetCommentByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var comment models.Comment
	if err := config.DB.First(&comment, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"erro": "Comentário não encontrado"})
	}
	return c.Status(200).JSON(fiber.Map{
		"id":        comment.ID,
		"content":   comment.Content,
		"taskId":    comment.TaskID,
		"createdAt": comment.CreatedAt,
		"updatedAt": comment.UpdatedAt,
	})
}

func UpdateComment(c *fiber.Ctx) error {
	id := c.Params("id")
	var comment models.Comment

	if err := config.DB.First(&comment, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"erro": "Comentário não encontrado"})
	}

	if err := c.BodyParser(&comment); err != nil {
		return c.Status(400).JSON(fiber.Map{"erro": "Entrada inválida"})
	}

	config.DB.Save(&comment)
	return c.Status(200).JSON(fiber.Map{
		"id":        comment.ID,
		"content":   comment.Content,
		"taskId":    comment.TaskID,
		"createdAt": comment.CreatedAt,
		"updatedAt": comment.UpdatedAt,
	})
}

func DeleteComment(c *fiber.Ctx) error {
	id := c.Params("id")

	result := config.DB.Where("id = ?", id).Delete(&models.Comment{})
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"erro": "Comentário não encontrado"})
	}

	return c.SendStatus(204)
}
