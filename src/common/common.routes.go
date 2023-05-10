package common

import (
	authV1 "github.com/bhumit070/hmm/src/apis/v1/auth"
	"github.com/bhumit070/hmm/src/constants"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	v1Routes := app.Group(constants.V1_PREFIX)

	authV1.RegisterAuthRoutes(v1Routes)

}
