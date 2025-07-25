package models

import (
	_ "gorm.io/gorm" // "import for side-effects" поставил _ перед импортом зачем?
	"time"
)

type User struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	Name       string     `json:"name" gorm:"not null"`
	Sirname    string     `json:"sirname" gorm:"not null"`
	Email      string     `json:"email" gorm:"unique;not null"`
	Activities []Activity `json:"activities" gorm:"foreignKey:UserID"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
