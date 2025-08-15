package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"size:190;uniqueIndex;not null"`
	Password  string `gorm:"not null"` // hash
	CreatedAt time.Time
	UpdatedAt time.Time
}
