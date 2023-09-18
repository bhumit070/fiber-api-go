package todoV1

import (
	authV1 "github.com/bhumit070/go_api_demo/src/apis/v1/auth"
	"github.com/bhumit070/go_api_demo/src/constants"
	"github.com/bhumit070/go_api_demo/src/db"
	"github.com/bhumit070/go_api_demo/src/db/models"
	"github.com/bhumit070/go_api_demo/src/helper"
	"github.com/bhumit070/go_api_demo/src/helper/request"
	"github.com/gofiber/fiber/v2"
)

func GetAllTodos(ctx *fiber.Ctx) error {
	var todos GetAllTodosResponse
	userInfo := ctx.Locals(constants.CONTEXT_USER_INFO_KEY).(authV1.SignupResponse)
	findTodoError := db.DB.Model(&models.TodoModel{}).Find(&todos, "user_id = ?", userInfo.ID).Error

	if findTodoError != nil {
		return helper.SendResponse(ctx, helper.Response{
			Message: "Error while fetching todos",
			Code:    500,
			Data:    nil,
		})
	}

	return helper.SendResponse(ctx, helper.Response{
		Message: "Todos",
		Code:    200,
		Data:    todos,
	})
}

func CreateTodo(ctx *fiber.Ctx) error {

	var body CreateTodoRequest
	shouldContinue := request.Validate(ctx, &body)

	if !shouldContinue {
		return nil
	}

	userInfo := ctx.Locals(constants.CONTEXT_USER_INFO_KEY).(authV1.SignupResponse)

	createdTodo := &models.TodoModel{
		Title:       body.Title,
		Description: body.Description,
		IsCompleted: body.IsCompleted,
		UserID:      userInfo.ID,
	}

	createTodoError := db.DB.Create(&createdTodo).Error

	if createTodoError != nil {
		return helper.SendResponse(ctx, helper.Response{
			Message: "Error while creating todo",
			Code:    500,
			Data:    nil,
		})
	}

	body.Id = createdTodo.ID
	body.UserID = createdTodo.UserID

	return helper.SendResponse(ctx, helper.Response{
		Message: "Todo created successfully",
		Code:    200,
		Data:    body,
	})

}
