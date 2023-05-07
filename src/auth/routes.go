package auth

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(versionRouter fiber.Router) {

	authRoutes := versionRouter.Group("/auth")
	authRoutes.Post("/login", Login)
}
