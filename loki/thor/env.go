package thor

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv(envs ...string) {
	godotenv.Load(envs...)
}

func GetEnv(key string, defaultValue ...string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return ""
}

func GetEnvInt(key string, defaultValue ...int) int {
	if value, ok := os.LookupEnv(key); ok {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return 0
}

func GetEnvBool(key string, defaultValue ...bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return false
}
