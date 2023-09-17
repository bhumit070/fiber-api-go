package constants

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT string

	V1_PREFIX            string = "/v1"
	V2_PREFIX            string = "/v2"
	SQL_DB_URL           string = ""
	JWT_SECRET           string = ""
	SOMETHING_WENT_WRONG string = "Something went wrong, please try again later!"
)

func InitEnvVariables() {
	godotenv.Load(".env")

	PORT = os.Getenv("PORT")
	SQL_DB_URL = os.Getenv("SQL_DB_URL")
	JWT_SECRET = os.Getenv("JWT_SECRET")
	fmt.Println("PORT IS ", PORT)
}
