package authV1

import (
	"github.com/bhumit070/go_api_demo/src/constants"
	"github.com/bhumit070/go_api_demo/src/db"
	"github.com/bhumit070/go_api_demo/src/db/models"
	"github.com/bhumit070/go_api_demo/src/helper"
	"github.com/gofiber/fiber/v2"
)

func TokenValidator(ctx *fiber.Ctx) error {

	tokenInfo, tokenValidationError := helper.VerifyJWT(ctx)

	if tokenValidationError != nil {
		return helper.SendResponse(ctx, helper.Response{
			Code:    401,
			Message: "UnAuthorized",
			Data:    error.Error(tokenValidationError),
		})
	}

	var response SignupResponse
	findUserError := db.DB.Model(&models.UserModel{}).First(&response, "id = ?", tokenInfo.ID).Error

	if findUserError != nil {
		return helper.SendResponse(ctx, helper.Response{
			Code:    401,
			Message: "UnAuthorized",
			Data:    error.Error(findUserError),
		})
	}

	// add custom value to Context
	ctx.Locals(constants.CONTEXT_USER_INFO_KEY, response)

	return ctx.Next()
}
