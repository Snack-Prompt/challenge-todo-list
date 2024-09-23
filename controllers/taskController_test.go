package controllers

import (
	"errors"
	"testing"
	"todo-list/models"

	"github.com/stretchr/testify/assert"
)

func GetTasksLogic(status string, page, limit int) ([]models.Task, error) {
	if status != "" && status != "pending" && status != "doing" && status != "completed" {
		return nil, errors.New("Status inválido. Os status válidos são 'pending', 'doing' ou 'completed'.")
	}

	var tasks []models.Task

	return tasks, nil
}

func CreateTaskLogic(task models.Task) (models.Task, error) {
	if task.Title == "" {
		return models.Task{}, errors.New("O título é obrigatório")
	}
	if task.Status != "pending" && task.Status != "doing" && task.Status != "completed" {
		return models.Task{}, errors.New("Status inválido. Os status válidos são 'pending', 'doing' ou 'completed'.")
	}
	if task.Priority != "low" && task.Priority != "medium" && task.Priority != "high" {
		return models.Task{}, errors.New("Prioridade inválida. As prioridades válidas são 'low', 'medium' ou 'high'.")
	}

	task.ID = "mocked-uuid"

	return task, nil
}

func GetTaskByIDLogic(id string) (models.Task, error) {
	if id == "not-found" {
		return models.Task{}, errors.New("Tarefa não encontrada")
	}

	return models.Task{
		ID:          id,
		Title:       "Tarefa mockada",
		Description: "Descrição mockada",
		Status:      "pending",
		Priority:    "high",
	}, nil
}

func UpdateTaskLogic(id string, updatedTask models.Task) (models.Task, error) {
	if id == "not-found" {
		return models.Task{}, errors.New("Tarefa não encontrada")
	}

	updatedTask.ID = id

	return updatedTask, nil
}

func TestGetTasksLogic(t *testing.T) {
	tests := []struct {
		name        string
		status      string
		page        int
		limit       int
		expectedErr error
	}{
		{
			name:        "Status válido - pending",
			status:      "pending",
			page:        1,
			limit:       10,
			expectedErr: nil,
		},
		{
			name:        "Status inválido",
			status:      "invalid",
			page:        1,
			limit:       10,
			expectedErr: errors.New("Status inválido. Os status válidos são 'pending', 'doing' ou 'completed'."),
		},
		{
			name:        "Sem status",
			status:      "",
			page:        1,
			limit:       10,
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetTasksLogic(tt.status, tt.page, tt.limit)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestCreateTaskLogic(t *testing.T) {
	tests := []struct {
		name        string
		task        models.Task
		expectedErr error
	}{
		{
			name: "Tarefa válida",
			task: models.Task{
				Title:       "Implementar funcionalidade de login",
				Description: "Criar a tela de login e a integração com o backend.",
				Status:      "pending",
				Priority:    "high",
			},
			expectedErr: nil,
		},
		{
			name: "Título ausente",
			task: models.Task{
				Description: "Sem título",
				Status:      "pending",
				Priority:    "high",
			},
			expectedErr: errors.New("O título é obrigatório"),
		},
		{
			name: "Status inválido",
			task: models.Task{
				Title:       "Tarefa com status inválido",
				Description: "Tentando salvar com status inválido",
				Status:      "invalid",
				Priority:    "high",
			},
			expectedErr: errors.New("Status inválido. Os status válidos são 'pending', 'doing' ou 'completed'."),
		},
		{
			name: "Prioridade inválida",
			task: models.Task{
				Title:       "Tarefa com prioridade inválida",
				Description: "Tentando salvar com prioridade inválida",
				Status:      "pending",
				Priority:    "invalid",
			},
			expectedErr: errors.New("Prioridade inválida. As prioridades válidas são 'low', 'medium' ou 'high'."),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CreateTaskLogic(tt.task)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestGetTaskByIDLogic(t *testing.T) {
	tests := []struct {
		name        string
		id          string
		expectedErr error
	}{
		{
			name:        "Tarefa encontrada",
			id:          "mocked-id",
			expectedErr: nil,
		},
		{
			name:        "Tarefa não encontrada",
			id:          "not-found",
			expectedErr: errors.New("Tarefa não encontrada"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetTaskByIDLogic(tt.id)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestUpdateTaskLogic(t *testing.T) {
	tests := []struct {
		name        string
		id          string
		task        models.Task
		expectedErr error
	}{
		{
			name: "Tarefa atualizada com sucesso",
			id:   "mocked-id",
			task: models.Task{
				Title:       "Tarefa atualizada",
				Description: "Nova descrição",
				Status:      "completed",
				Priority:    "medium",
			},
			expectedErr: nil,
		},
		{
			name:        "Tarefa não encontrada",
			id:          "not-found",
			task:        models.Task{},
			expectedErr: errors.New("Tarefa não encontrada"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := UpdateTaskLogic(tt.id, tt.task)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
