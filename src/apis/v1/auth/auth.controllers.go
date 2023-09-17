package authV1

import (
	"fmt"
	"strings"
	"time"

	"github.com/bhumit070/go_api_demo/src/constants"
	"github.com/bhumit070/go_api_demo/src/db"
	"github.com/bhumit070/go_api_demo/src/db/models"
	"github.com/bhumit070/go_api_demo/src/helper"
	"github.com/bhumit070/go_api_demo/src/helper/request"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type LoginBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `password:"password" validate:"required,min=6"`
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

type LoginUser struct {
	Password string `json:"password"`
	SignupResponse
}

type LoginResponse struct {
	SignupResponse
	Token string `json:"token"`
}

func Login(ctx *fiber.Ctx) error {
	var body LoginBody

	shouldContinue := request.Validate(ctx, &body)

	if !shouldContinue {
		return nil
	}

	var existingUser LoginUser
	findingExistingUserError := db.DB.Model(&models.UserModel{}).First(&existingUser, "email = ?", body.Email).Error

	if findingExistingUserError != nil {
		return helper.SendResponse(ctx, helper.Response{
			Code:    400,
			Message: "Incorrect credentials provided!",
			Data:    nil,
		})
	}

	comparePasswordError := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(body.Password))

	if comparePasswordError != nil {
		return helper.SendResponse(ctx, helper.Response{
			Code:    400,
			Message: "Incorrect credentials provided!",
			Data:    nil,
		})
	}

	var response LoginResponse
	response.SignupResponse = existingUser.SignupResponse

	token, generatingTokenError := helper.GenerateJwt(response.ID)

	if generatingTokenError != nil {
		return helper.SendResponse(ctx, helper.Response{
			Code:    500,
			Message: constants.SOMETHING_WENT_WRONG,
		})
	}

	response.Token = token

	return helper.SendResponse(ctx, helper.Response{
		Code:    200,
		Message: "Login Successful!",
		Data:    response,
	})
}

func Register(ctx *fiber.Ctx) error {
	var body SignupBody

	shouldContinue := request.Validate(ctx, &body)

	if !shouldContinue {
		return nil
	}

	var existingUser SignupResponse
	findingExistingUserError := db.DB.Model(&models.UserModel{}).First(&existingUser, "email = ?", body.Email).Error

	if findingExistingUserError != nil && findingExistingUserError.Error() != "record not found" {
		return helper.SendResponse(ctx, helper.Response{
			Code:    400,
			Message: constants.SOMETHING_WENT_WRONG,
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

func ValidateToken(ctx *fiber.Ctx) error {
	token := strings.Trim(strings.Split(string(ctx.Request().Header.Peek("Authorization")), "Bearer")[1], " ")

	tokenInfo, _ := helper.VerifyJWT(token)

	return helper.SendResponse(ctx, helper.Response{
		Code:    200,
		Message: "Token Validated!",
		Data:    tokenInfo,
	})
}
