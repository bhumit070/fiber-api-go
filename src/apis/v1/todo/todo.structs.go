package todoV1

import (
	"github.com/bhumit070/go_api_demo/src/db/models"
)

type GetAllTodosResponse struct {
	ID          uint              `gorm:"primaryKey" json:"id"`
	Title       string            `gorm:"not null" json:"name"`
	Description string            `gorm:"not null" json:"description"`
	IsCompleted bool              `gorm:"default:false" json:"isCompleted"`
	UserID      uint              `gorm:"not null" json:"userId"`
	UserInfo    *models.UserModel `gorm:"foreignKey:UserID" json:"userInfo"`
}

type CreateTodoRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsCompleted bool   `json:"isCompleted" default:"false"`
	UserID      uint   `gorm:"not null" json:"userId"`
	Id          uint   `json:"id"`
}
