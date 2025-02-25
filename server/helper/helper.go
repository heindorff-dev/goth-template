package helper

import (
	"errors"
	"os"
)

func MustGetEnv(key string) (string, error) {
	envValue := os.Getenv(key)
	if envValue == "" {
		errMessage := "The environment variable '" + key + "' doesn't exist or is not set."
		return "", errors.New(errMessage)
	}
	return envValue, nil
}
