package models

import "time"

type Task struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Titulo     string    `gorm:"type:nvarchar(255);not null" json:"titulo" binding:"required"`
	Completada bool      `gorm:"not null;default:false" json:"completada"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
