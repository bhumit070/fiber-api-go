package helper

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Code    int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(ctx *fiber.Ctx, response Response) error {
	ctx.Status(response.Code)
	return ctx.JSON(response)
}
