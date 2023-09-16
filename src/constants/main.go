package constants

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT string

	V1_PREFIX  string = "/v1"
	V2_PREFIX  string = "/v2"
	SQL_DB_URL string
)

func InitEnvVariables() {
	godotenv.Load(".env")

	PORT = os.Getenv("PORT")
	SQL_DB_URL = os.Getenv("SQL_DB_URL")
	fmt.Println("PORT IS ", PORT)
}
