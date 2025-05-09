package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Envirounment struct {
	DbDatabase string
	DbPassword string
	DbPort     int
	DbServer   string
	DbUser     string
	ServerPort int
	CorsOrigin string
}

var Env = newEnvirounment()

func newEnvirounment() *Envirounment {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Envirounment{
		DbDatabase: os.Getenv("DB_DATABASE"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbPort:     getEnvirounmentVariableAsInt("DB_PORT"),
		DbServer:   os.Getenv("DB_SERVER"),
		DbUser:     os.Getenv("DB_USER"),
		ServerPort: getEnvirounmentVariableAsInt("SERVER_PORT"),
		CorsOrigin: os.Getenv("CORS_ORIGIN"),
	}
}

func getEnvirounmentVariableAsInt(key string) int {
	value, err := strconv.Atoi(os.Getenv(key))

	if err != nil {
		return 0
	}

	return value
}
