package todoV1

import (
	authV1 "github.com/bhumit070/go_api_demo/src/apis/v1/auth"
	"github.com/bhumit070/go_api_demo/src/constants"
	"github.com/bhumit070/go_api_demo/src/db"
	"github.com/bhumit070/go_api_demo/src/db/models"
	"github.com/bhumit070/go_api_demo/src/helper"
	"github.com/bhumit070/go_api_demo/src/helper/request"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllTodos(ctx *fiber.Ctx) error {
	var todos []GetAllTodosResponse
	userInfo := ctx.Locals(constants.CONTEXT_USER_INFO_KEY).(authV1.SignupResponse)
	findTodoError := db.DB.Model(&models.TodoModel{}).
		Where("user_id = ?", userInfo.ID).
		Find(&todos).Error

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

func GetOneTodo(ctx *fiber.Ctx) error {
	todoId := ctx.Params("todoId")

	if todoId == "" {
		return helper.SendResponse(ctx, helper.Response{
			Message: "Todo id is required",
			Code:    400,
			Data:    nil,
		})
	}

	var todo GetOneTodosResponse
	findTodoError := db.DB.Model(&models.TodoModel{}).
		Preload("UserInfo", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name", "email")
		}).
		Find(&todo, "id = ?", todoId).Error

	if findTodoError != nil {
		return helper.SendResponse(ctx, helper.Response{
			Message: findTodoError.Error(),
			Code:    404,
			Data:    nil,
		})
	}

	return helper.SendResponse(ctx, helper.Response{
		Message: "Todo",
		Code:    200,
		Data:    todo,
	})
}

func DeleteTodo(ctx *fiber.Ctx) error {
	todoId := ctx.Params("todoId")

	if todoId == "" {
		return helper.SendResponse(ctx, helper.Response{
			Message: "Todo id is required",
			Code:    400,
			Data:    nil,
		})
	}

	db.DB.Delete(&models.TodoModel{}, "id = ?", todoId)

	return helper.SendResponse(ctx, helper.Response{
		Message: "Todo deleted successfully",
		Code:    200,
		Data:    nil,
	})
}
