package main

import (
	"fmt"
	"log"

	"github.com/bhumit070/go_api_demo/src/apis/v1/auth"
	"github.com/bhumit070/go_api_demo/src/common"
	"github.com/bhumit070/go_api_demo/src/constants"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	constants.InitEnvVariables()

	app := fiber.New()
	app.Use(logger.New())

	common.RegisterRoutes(app)

	app.All("*", func(ctx *fiber.Ctx) error {
		ctx.Status(404)
		return ctx.JSON(auth.Response{
			Status:  404,
			Message: "Route Not Found!",
		})
	})

	log.Fatal(app.Listen(":" + constants.PORT))

	fmt.Println("Server is running on port", constants.PORT)

}
