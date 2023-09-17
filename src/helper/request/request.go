package request

import (
	"strings"

	"github.com/bhumit070/go_api_demo/src/helper"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// write a genetic function
func Validate(ctx *fiber.Ctx, body interface{}) bool {
	err := ctx.BodyParser(&body)

	if err != nil {
		helper.SendResponse(ctx, helper.Response{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return false
	}

	validate := validator.New()
	err = validate.Struct(body)

	if err != nil {
		var validationErrors = make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errorReason := err.ActualTag()
			if err.ActualTag() == strings.ToLower(err.Field()) {
				errorReason = "invalid"
			}
			validationErrors[err.Field()] = err.Field() + " is " + errorReason

		}

		helper.SendResponse(ctx, helper.Response{
			Code:    400,
			Message: "Validation Error!",
			Data:    validationErrors,
		})
		return false
	}
	return true
}
