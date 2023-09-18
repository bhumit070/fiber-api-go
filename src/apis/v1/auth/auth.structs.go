package authV1

import "time"

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
