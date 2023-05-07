package common

import (
	"github.com/bhumit070/hmm/src/auth"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	v1Routes := app.Group("/v1")

	auth.RegisterAuthRoutes(v1Routes)

}
