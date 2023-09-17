package main

import (
	"fmt"
	"log"

	"github.com/bhumit070/go_api_demo/src/apis"
	"github.com/bhumit070/go_api_demo/src/constants"
	"github.com/bhumit070/go_api_demo/src/db"
	"github.com/bhumit070/go_api_demo/src/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	helper.VerifyJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NjUsImV4cCI6MTY5NTU0MjYwOSwibmJmIjoxNjk0OTM3ODA5LCJpYXQiOjE2OTQ5Mzc4MDl9.zw8BrqXst6AQvKvcwiD1LYZztdu5ZPn3A3xv_Hpl47k")

	constants.InitEnvVariables()
	db.InitDB()
	app := fiber.New()
	app.Use(logger.New())

	apis.RegisterRoutes(app)

	log.Fatal(app.Listen(":" + constants.PORT))

	fmt.Println("Server is running on port", constants.PORT)
}
