package envutil

import (
	"os"
	"strconv"
	"strings"
)

// GetEnv ...
var GetEnv = func(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

// GetEnvAsInt ...
func GetEnvAsInt(key string, defaultVal int) int {
	valueStr := GetEnv(key, "")
	if val, err := strconv.Atoi(valueStr); err == nil {
		return val
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
	return strings.Split(valStr, sep)
}

// GetEnvAsMapBool ...
func GetEnvAsMapBool(key string, defaultVal map[string]bool, sep string) map[string]bool {
	valStr := GetEnv(key, "")
	if valStr == "" {
		return defaultVal
	}
	vals := strings.Split(valStr, sep)
	valMap := make(map[string]bool, len(vals))
	for i := range vals {
		valMap[vals[i]] = true
	}
	return valMap
}
