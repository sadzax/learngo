package models

import (
	_ "gorm.io/gorm" // "import for side-effects" поставил _ перед импортом зачем?
	"time"
)

type Activity struct {
	ID           uint          `json:"id" gorm:"primaryKey"`
	Name         string        `json:"name" gorm:"not null"`
	Description  string        `json:"description"`
	UserID       uint          `json:"user_id" gorm:"not null"`
	User         User          `json:"user" gorm:"foreignKey:UserID"`
	ActivityLogs []ActivityLog `json:"activity_logs" gorm:"foreignKey:ActivityID"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}
