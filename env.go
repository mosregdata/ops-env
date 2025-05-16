package opsenv

import (
	"fmt"
	"os"
	"strings"
)

// HasEnv проверяет существование переменной.
func HasEnv(key string) bool {
	_, exists := os.LookupEnv(key)
	return exists
}

// GetEnv возвращает значение переменной окружения или ошибку.
func GetEnv(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return "", fmt.Errorf("environment variable %s not found", key)
	}
	return value, nil
}

// GetEnvDefault возвращает значение переменной или дефолтное.
func GetEnvDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// SetEnv устанавливает переменную окружения.
func SetEnv(key, value string) error {
	return os.Setenv(key, value)
}

// UnsetEnv удаляет переменную окружения.
func UnsetEnv(key string) error {
	return os.Unsetenv(key)
}

// GetAllEnv возвращает все переменные окружения.
func GetAllEnv() map[string]string {
	result := make(map[string]string)
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		if len(pair) == 2 {
			result[pair[0]] = pair[1]
		}
	}
	return result
}
