package constants

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT                  string = "8080"
	V1_PREFIX             string = "/v1"
	V2_PREFIX             string = "/v2"
	SQL_DB_URL            string = ""
	JWT_SECRET            string = "something_super_secret"
	SOMETHING_WENT_WRONG  string = "Something went wrong, please try again later!"
	CONTEXT_USER_INFO_KEY string = "USER_DETAILS"
)

func InitEnvVariables() {
	godotenv.Load(".env")

	PORT = os.Getenv("PORT")
	SQL_DB_URL = os.Getenv("SQL_DB_URL")

	if os.Getenv("JWT_SECRET") != "" {
		JWT_SECRET = os.Getenv("JWT_SECRET")
	}
}
