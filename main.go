package main

import (
	"fmt"
	"log"

	"github.com/bhumit070/hmm/src/common"
	"github.com/bhumit070/hmm/src/constants"
	"github.com/gofiber/fiber/v2"
)

func main() {
	constants.InitEnvVariables()

	app := fiber.New()

	common.RegisterRoutes(app)

	log.Fatal(app.Listen(constants.PORT))

	fmt.Println("Server is running on port", constants.PORT)

}
