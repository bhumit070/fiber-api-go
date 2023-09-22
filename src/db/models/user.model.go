package models

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	ID        uint         `gorm:"primaryKey" json:"id,omitempty"`
	Name      string       `gorm:"not null" json:"name,omitempty"`
	Email     string       `gorm:"unique" json:"email,omitempty"`
	Password  string       `gorm:"not null" json:"password,omitempty"`
	CreatedAt time.Time    `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time    `gorm:"autoCreateTime" json:"-"`
	DeletedAt time.Time    `gorm:"default:null" json:"-"`
	Todos     *[]TodoModel `gorm:"foreignKey:UserID" json:"todos,omitempty"`
}

func (UserModel) TableName() string {
	return "users"
}

func (user *UserModel) AfterFind(tx *gorm.DB) (err error) {
	return
}
