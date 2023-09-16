package main

import (
	"fmt"
	"log"

	"github.com/bhumit070/go_api_demo/src/apis"
	"github.com/bhumit070/go_api_demo/src/constants"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	constants.InitEnvVariables()

	app := fiber.New()
	app.Use(logger.New())

	apis.RegisterRoutes(app)

	log.Fatal(app.Listen(":" + constants.PORT))

	fmt.Println("Server is running on port", constants.PORT)

}
