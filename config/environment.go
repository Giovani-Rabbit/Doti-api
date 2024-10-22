package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Envirounment struct {
	DbServer   string
	DbUser     string
	DbPassword string
	DbPort     int
	DbDatabase string
}

var Env = newEnvirounment()

func newEnvirounment() *Envirounment {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Envirounment{
		DbServer:   os.Getenv("DB_SERVER"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbDatabase: os.Getenv("DB_DATABASE"),
		DbPort:     getEnvirounmentVariableAsInt("DB_PORT"),
	}
}

func getEnvirounmentVariableAsInt(key string) int {
	value, err := strconv.Atoi(os.Getenv(key))

	if err != nil {
		return 0
	}

	return value
}
