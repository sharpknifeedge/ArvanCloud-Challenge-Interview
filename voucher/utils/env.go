package utils

import (
	"os"
	"strconv"
)

func Env(key, defaultVal string) string {
	s := os.Getenv(key)
	if len(s) == 0 {
		return defaultVal
	}

	return s
}

func EnvInt(key string, defaultVal int) int {
	s := os.Getenv(key)
	if len(s) == 0 {
		return defaultVal
	}

	val, err := strconv.Atoi(s)
	if err != nil {
		return defaultVal
	}

	return val
}
