package db

import (
	"fmt"

	"github.com/bhumit070/go_api_demo/src/constants"
	"github.com/bhumit070/go_api_demo/src/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB = nil

func InitDB() {
	db, err := gorm.Open(postgres.Open(constants.SQL_DB_URL), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	} else {
		DB = db
		fmt.Println("Connected to database!")
		RunMigration()
	}

}

func RunMigration() {
	if DB == nil {
		return
	}
	err := DB.AutoMigrate(&models.UserModel{})
	if err != nil {
		panic("Failed to migrate!")
	} else {
		fmt.Println("Migration successful!")
	}
}
