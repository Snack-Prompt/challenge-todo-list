package models

import "time"

type Task struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"not null" validate:"required,min=3,max=100" json:"title"`
	Description string    `gorm:"not null" validate:"required,min=3,max=500" json:"description"`
	Status      string    `gorm:"not null" validate:"required,oneof=pending doing completed" json:"status"`
	Priority    string    `gorm:"not null;default:'medium'" validate:"required,oneof=low medium high" json:"priority"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
