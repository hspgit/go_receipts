package models

import (
	"time"
)

// User model for database operations
type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"column:name;not null"`
	Email     string    `gorm:"column:email;unique;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}
