package models

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoCreateTime" json:"updatedAt"`
	DeletedAt time.Time `gorm:"default:null" json:"deletedAt,omitempty"`
}

func (UserModel) TableName() string {
	return "users"
}

func (user *UserModel) AfterFind(tx *gorm.DB) (err error) {
	return
}
