package main

import (
	"fmt"
	"log"

	"github.com/bhumit070/go_api_demo/src/common"
	"github.com/bhumit070/go_api_demo/src/constants"
	"github.com/gofiber/fiber/v2"
)

func main() {
	constants.InitEnvVariables()

	app := fiber.New()

	common.RegisterRoutes(app)

	log.Fatal(app.Listen(":" + constants.PORT))

	fmt.Println("Server is running on port", constants.PORT)

}
