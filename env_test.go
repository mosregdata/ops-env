package opsenv

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	key := "TEST_KEY"
	value := "test_value"
	os.Setenv(key, value)
	defer os.Unsetenv(key)

	result, err := GetEnv(key)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if result != value {
		t.Errorf("expected %s, got %s", value, result)
	}

	_, err = GetEnv("NON_EXISTENT_KEY")
	if err == nil {
		t.Error("expected error for non-existent key, got nil")
	}
}

func TestSetEnv(t *testing.T) {
	key := "TEST_SET_KEY"
	value := "test_set_value"

	err := SetEnv(key, value)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	result, exists := os.LookupEnv(key)
	if !exists {
		t.Error("expected key to exist")
	}
	if result != value {
		t.Errorf("expected %s, got %s", value, result)
	}

	os.Unsetenv(key)
}

func TestUnsetEnv(t *testing.T) {
	key := "TEST_UNSET_KEY"
	os.Setenv(key, "some_value")

	err := UnsetEnv(key)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	_, exists := os.LookupEnv(key)
	if exists {
		t.Error("expected key to be unset")
	}
}

func TestGetAllEnv(t *testing.T) {
	os.Setenv("KEY1", "value1")
	os.Setenv("KEY2", "value2")
	defer os.Unsetenv("KEY1")
	defer os.Unsetenv("KEY2")

	result := GetAllEnv()
	if result["KEY1"] != "value1" {
		t.Errorf("expected KEY1=value1, got %s", result["KEY1"])
	}
	if result["KEY2"] != "value2" {
		t.Errorf("expected KEY2=value2, got %s", result["KEY2"])
	}
}

func TestHasEnv(t *testing.T) {
	key := "TEST_HAS_KEY"
	os.Setenv(key, "some_value")
	defer os.Unsetenv(key)

	if !HasEnv(key) {
		t.Error("expected true for existing key")
	}

	if HasEnv("NON_EXISTENT_KEY") {
		t.Error("expected false for non-existent key")
	}
}

func TestGetEnvDefault(t *testing.T) {
	key := "TEST_DEFAULT_KEY"
	value := "test_default_value"
	defaultValue := "default"

	os.Setenv(key, value)
	defer os.Unsetenv(key)
	result := GetEnvDefault(key, defaultValue)
	if result != value {
		t.Errorf("expected %s, got %s", value, result)
	}

	// Тест с несуществующей переменной
	result = GetEnvDefault("NON_EXISTENT_KEY", defaultValue)
	if result != defaultValue {
		t.Errorf("expected %s, got %s", defaultValue, result)
	}
}
