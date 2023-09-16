package auth

import (
	"fmt"

	"github.com/bhumit070/go_api_demo/src/helper"
	"github.com/gofiber/fiber/v2"
)

type LoginBody struct {
	Email      string `json:"email"`
	Password   string `password:"password"`
	ApiVersion string `json:"apiVersion"`
}

func Login(ctx *fiber.Ctx) error {
	var body LoginBody
	err := ctx.BodyParser(&body)

	if err != nil {
		fmt.Println(err)
		return helper.SendResponse(
			ctx,
			helper.Response{
				Code:    400,
				Data:    nil,
				Message: err.Error(),
			},
		)
	}

	return helper.SendResponse(ctx, helper.Response{
		Code:    200,
		Data:    body,
		Message: "Login Successful!",
	})
}
