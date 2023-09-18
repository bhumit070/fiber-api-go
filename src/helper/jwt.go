package helper

import (
	"errors"
	"strings"
	"time"

	"github.com/bhumit070/go_api_demo/src/constants"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type CustomJwtClaims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}

var JWT_SECRET []byte = []byte(constants.JWT_SECRET)

func GenerateJwt(userId uint) (string, error) {
	claims := &CustomJwtClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWT_SECRET)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyJWT(ctx *fiber.Ctx) (*CustomJwtClaims, error) {

	tokenString := string(ctx.Request().Header.Peek("Authorization"))

	if strings.HasPrefix(tokenString, "Bearer") {
		splitString := strings.Split(tokenString, "Bearer ")
		if len(splitString) >= 2 {
			tokenString = strings.Trim(splitString[1], "")
		}
	} else {
		tokenString = strings.Trim(tokenString, "")
	}

	claims := &CustomJwtClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return JWT_SECRET, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token has been passed")
	}

	return claims, nil
}
