package models

import (
	_ "gorm.io/gorm" // "import for side-effects" поставил _ перед импортом зачем?
	"time"
)

type ActivityLog struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	ActivityID uint      `json:"activity_id" gorm:"not null"`
	Activity   Activity  `json:"activity" gorm:"foreignKey:ActivityID"`
	Hours      float64   `json:"hours" gorm:"not null"`
	LoggedAt   time.Time `json:"logged_at" gorm:"default:CURRENT_TIMESTAMP"`
	CreatedAt  time.Time `json:"created_at"`
}
