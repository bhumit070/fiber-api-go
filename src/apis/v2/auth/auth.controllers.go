package authV2

import (
	"github.com/bhumit070/go_api_demo/src/constants"
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
		return helper.SendResponse(ctx, helper.Response{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
	}

	body.ApiVersion = constants.V2_PREFIX

	return helper.SendResponse(ctx, helper.Response{
		Code:    200,
		Message: "Login Successful!",
		Data:    body,
	})
}
