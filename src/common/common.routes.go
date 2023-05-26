package common

import (
	authV1 "github.com/bhumit070/go_api_demo/src/apis/v1/auth"
	authV2 "github.com/bhumit070/go_api_demo/src/apis/v2/auth"
	"github.com/bhumit070/go_api_demo/src/constants"
	"github.com/gofiber/fiber/v2"
)

func registerV1Routes(apiRouter fiber.Router) {
	v1Routes := apiRouter.Group(constants.V1_PREFIX)

	authV1.RegisterAuthRoutes(v1Routes)
}

func registerV2Routes(apiRouter fiber.Router) {
	v2Routes := apiRouter.Group(constants.V2_PREFIX)

	authV2.RegisterAuthRoutes(v2Routes)
}

func RegisterRoutes(app *fiber.App) {

	apiRouter := app.Group("/api")

	registerV1Routes(apiRouter)
	registerV2Routes(apiRouter)
}
