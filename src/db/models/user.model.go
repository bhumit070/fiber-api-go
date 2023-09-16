package models

import (
	"time"
)

type UserModel struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string    `gorm:"unique"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
	DeletedAt time.Time
}
