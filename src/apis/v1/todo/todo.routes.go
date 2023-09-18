package todoV1

import (
	authV1 "github.com/bhumit070/go_api_demo/src/apis/v1/auth"
	"github.com/gofiber/fiber/v2"
)

func RegisterTodoRoutes(versionRouter fiber.Router) {
	todoRouter := versionRouter.Group("/todos")

	todoRouter.Use(authV1.TokenValidator)

	todoRouter.
		Get("/", GetAllTodos).
		Get("/:todoId", GetOneTodo).
		Post("/", CreateTodo).
		Delete("/:todoId", DeleteTodo)
}
