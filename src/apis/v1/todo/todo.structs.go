package todoV1

import (
	"github.com/bhumit070/go_api_demo/src/db/models"
)

type GetAllTodosResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"name"`
	Description string `json:"description"`
	IsCompleted bool   `json:"isCompleted"`
}

type GetOneTodosResponse struct {
	ID          uint             `json:"id"`
	Title       string           `json:"name"`
	Description string           `json:"description"`
	IsCompleted bool             `json:"isCompleted"`
	UserID      uint             `gorm:"foreignKey:ID" json:"-"`
	UserInfo    models.UserModel `json:"userInfo" gorm:"foreignKey:UserID"`
}

type CreateTodoRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsCompleted bool   `json:"isCompleted" default:"false"`
	UserID      uint   `gorm:"not null" json:"userId"`
	Id          uint   `json:"id"`
}
