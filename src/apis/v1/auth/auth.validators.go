package authV1

import (
	"strings"

	"github.com/bhumit070/go_api_demo/src/constants"
	"github.com/bhumit070/go_api_demo/src/db"
	"github.com/bhumit070/go_api_demo/src/db/models"
	"github.com/bhumit070/go_api_demo/src/helper"
	"github.com/gofiber/fiber/v2"
)

func TokenValidator(ctx *fiber.Ctx) error {

	token := string(ctx.Request().Header.Peek("Authorization"))

	if strings.HasPrefix(token, "Bearer") {
		splitString := strings.Split(token, "Bearer")
		if len(splitString) >= 2 {
			token = strings.Trim(splitString[1], "")
		}
	}

	tokenInfo, tokenValidationError := helper.VerifyJWT(token)

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
