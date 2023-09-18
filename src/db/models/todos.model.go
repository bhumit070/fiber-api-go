package models

import (
	"time"

	"gorm.io/gorm"
)

type TodoModel struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"not null" json:"name"`
	Description string    `gorm:"not null" json:"description"`
	IsCompleted bool      `gorm:"default:false" json:"isCompleted"`
	UserID      uint      `gorm:"not null" json:"userId"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"autoCreateTime" json:"updatedAt"`
	DeletedAt   time.Time `gorm:"default:null" json:"deletedAt,omitempty"`
	UserInfo    UserModel `gorm:"foreignKey:UserID" json:"userInfo,omitempty"`
}

type GetAllTodosResponse struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
	IsCompleted bool   `gorm:"default:false" json:"isCompleted"`
	UserID      uint   `gorm:"not null" json:"userId"`
}

type CreateTodoRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsCompleted bool   `json:"isCompleted" default:"false"`
	UserID      uint   `gorm:"not null" json:"userId"`
	Id          uint   `json:"id"`
}

func (TodoModel) TableName() string {
	return "todos"
}
