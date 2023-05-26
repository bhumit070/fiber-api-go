package auth

import (
	"github.com/bhumit070/go_api_demo/src/constants"
	"github.com/gofiber/fiber/v2"
)

type LoginBody struct {
	Email      string `json:"email"`
	Password   string `password:"password"`
	ApiVersion string `json:"apiVersion"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func Login(ctx *fiber.Ctx) error {
	var body LoginBody
	err := ctx.BodyParser(&body)

	if err != nil {
		return ctx.JSON(Response{
			Status:  400,
			Message: "Error!",
		})
	}

	body.ApiVersion = constants.V2_PREFIX

	return ctx.JSON(body)
}
