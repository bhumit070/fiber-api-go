package constants

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT string
)

func InitEnvVariables() {
	godotenv.Load(".env")

	PORT = os.Getenv("PORT")

}
