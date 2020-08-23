package envutil

import (
	"os"
	"strconv"
	"strings"
)

// GetEnv ...
func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// GetEnvAsInt ...
func GetEnvAsInt(key string, defaultVal int) int {
	valueStr := GetEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}

// GetEnvAsBool ...
func GetEnvAsBool(key string, defaultVal bool) bool {
	valStr := GetEnv(key, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}
	return defaultVal
}

// GetEnvAsSlice ...
func GetEnvAsSlice(key string, defaultVal []string, sep string) []string {
	valStr := GetEnv(key, "")
	if valStr == "" {
		return defaultVal
	}
	val := strings.Split(valStr, sep)
	return val
}
