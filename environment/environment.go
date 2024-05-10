package environment

import (
	"github.com/joho/godotenv"
	"os"
	"log"
	"fmt"
)

var(
	APP_PORT string
	ENVIRONMENT string
	DBHost string
	DBUser string
	DBPassword string
	DBName string
	DBPort string
	DBSSLMode string
	DBTimezone string
	Conn string
)

func Config(){

APP_PORT = os.Getenv("PORT")
ENVIRONMENT = os.Getenv("ENVIRONMENT")
DBHost = os.Getenv("DBHost")
DBUser = os.Getenv("DBUser")
DBPassword = "Pass#1230"//os.Getenv("DBPassword")
DBName = os.Getenv("DBName")
DBPort = os.Getenv("DBPort")
DBSSLMode = os.Getenv("DBSSLMode")
DBTimezone = os.Getenv("DBTimezone")
Conn=fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",DBHost, DBUser, DBPassword, DBName, DBPort, DBSSLMode, DBTimezone)

}


func LoadEnvironmentVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Unable to load environment file.")
	}
	Config()
}