package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/bhumit070/go_api_demo/src/db"
	"github.com/bhumit070/go_api_demo/src/db/models"
	"github.com/bhumit070/go_api_demo/src/helper"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type LoginBody struct {
	Email      string `json:"email" validate:"required,email"`
	Password   string `password:"password"`
	ApiVersion string `json:"apiVersion"`
}

type SignupBody struct {
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Password string `password:"password" validate:"required,min=6"`
}

type SignupResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func Login(ctx *fiber.Ctx) error {
	var body LoginBody
	err := ctx.BodyParser(&body)

	if err != nil {
		fmt.Println(err)
		return helper.SendResponse(
			ctx,
			helper.Response{
				Code:    400,
				Data:    nil,
				Message: err.Error(),
			},
		)
	}

	return helper.SendResponse(ctx, helper.Response{
		Code:    200,
		Data:    body,
		Message: "Login Successful!",
	})
}

func Register(ctx *fiber.Ctx) error {
	var body SignupBody
	err := ctx.BodyParser(&body)

	if err != nil {
		return helper.SendResponse(ctx, helper.Response{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
	}

	validate := validator.New()
	err = validate.Struct(body)

	if err != nil {
		var validationErrors = make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errorReason := err.ActualTag()
			if err.ActualTag() == strings.ToLower(err.Field()) {
				errorReason = "invalid"
			}
			validationErrors[err.Field()] = err.Field() + " is " + errorReason

		}

		return helper.SendResponse(ctx, helper.Response{
			Code:    400,
			Message: "Validation Error!",
			Data:    validationErrors,
		})
	}

	var existingUser SignupResponse
	findingExistingUserError := db.DB.Model(&models.UserModel{}).First(&existingUser, "email = ?", body.Email).Error

	if findingExistingUserError != nil && findingExistingUserError.Error() != "record not found" {
		return helper.SendResponse(ctx, helper.Response{
			Code:    400,
			Message: "Something went wrong, please try again later!",
			Data:    nil,
		})
	}

	if existingUser.Email == body.Email {
		return helper.SendResponse(ctx, helper.Response{
			Code:    400,
			Message: "Email already exists!",
			Data:    nil,
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		fmt.Println("Error hashing password:", err)
	}

	body.Password = string(hashedPassword)

	// get saved user or error
	result := db.DB.Create(&models.UserModel{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	})

	if result.Error != nil {
		return helper.SendResponse(ctx, helper.Response{
			Code:    400,
			Message: "Signup Failed!",
			Data:    result.Error,
		})
	}
	var user SignupResponse
	findingUserError := db.DB.Model(&models.UserModel{}).First(&user).Error

	if findingUserError != nil {
		user = SignupResponse{
			Email: body.Email,
			Name:  body.Name,
		}
	}

	return helper.SendResponse(ctx, helper.Response{
		Code:    200,
		Message: "Signup Successful!",
		Data:    user,
	})
}
