package authV1

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(versionRouter fiber.Router) {

	authRoutes := versionRouter.Group("/auth")
	authRoutes.Post("/login", Login)
	authRoutes.Post("/register", Register)
	authRoutes.Get("validate-token", ValidateToken)
}
