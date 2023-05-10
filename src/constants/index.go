package constants

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT      string
	V1_PREFIX string = "/v1"
)

func InitEnvVariables() {
	godotenv.Load(".env")

	PORT = os.Getenv("PORT")

}
