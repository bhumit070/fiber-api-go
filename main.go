package main

import (
	"fmt"
	"log"

	"github.com/bhumit070/hmm/src/common"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	common.RegisterRoutes(app)

	const PORT = ":4040"

	log.Fatal(app.Listen(PORT))

	fmt.Println("Server is running on port", PORT)

}
