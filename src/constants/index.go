package constants

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT string

	V1_PREFIX string = "/v1"
	V2_PREFIX string = "/v2"
)

func InitEnvVariables() {
	godotenv.Load(".env")

	PORT = os.Getenv("PORT")
	fmt.Println("PORT IS ", PORT)
}
