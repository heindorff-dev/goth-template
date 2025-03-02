package helper

import (
	"os"

	"github.com/joho/godotenv"
)

func MustGetEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	envValue := os.Getenv(key)
	if envValue == "" {
		errMessage := "The environment variable '" + key + "' doesn't exist or is not set."
		panic(errMessage)
	}
	return envValue
}
