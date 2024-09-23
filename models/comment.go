package models

import "time"

type Comment struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Content   string    `gorm:"not null" json:"content"`
	TaskID    string    `gorm:"not null" json:"taskId"`
	Task      Task      `gorm:"foreignKey:TaskID"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
