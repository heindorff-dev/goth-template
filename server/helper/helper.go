package helper

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) (string, error) {
	envValue := os.Getenv(key)
	if envValue == "" {
		errMessage := "The environment variable '" + key + "' doesn't exist or is not set."
		return "", errors.New(errMessage)
	}
	return envValue, nil
}

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
