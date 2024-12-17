package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique"`
	Password  string    `json:"password" gorm:"not null"`
	Role      string    `json:"role" gorm:"not null"`
	IsDeleted bool      `json:"is_deleted" gorm:"default:false"`  // Флаг удаления
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"` // Автоматическая метка времени
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"` // Автоматическая метка времени
}

type SwagUser struct {
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password" gorm:"not null"`
}

type SignInInput struct {
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password" gorm:"not null"`
}
