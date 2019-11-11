package env

import (
	"log"
	"os"
	"strconv"
)

// GetEnv will get the env with the key envKey and return its value as a string
// If no value is found it will panic
func GetEnv(envKey string, errorMsg string) string {
	value, exists := os.LookupEnv(envKey)
	if !exists {
		log.Fatalf("'%s' %s", envKey, errorMsg)
	}
	return value
}

// GetIntOptional will return the value of the key as a int, or use default if no value is found
func GetIntOptional(key string, defaultValue int) int {
	env, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("Key '%s' was not found, using default value %d", key, defaultValue)
		return defaultValue
	}
	number, err := strconv.Atoi(env)
	if err != nil {
		panic(err)
	}

	return number
}

// GetStringOptional will return the value of the key as a string, or use default if no value is found
func GetStringOptional(key string, defaultValue string) string {
	env, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("Key '%s' was not found, using default value %s", key, defaultValue)
		return defaultValue
	}
	return env
}

// GetBoolOptional will return the value of the key as a bool, or use default if no value is found
func GetBoolOptional(key string, defaultValue bool) bool {
	env, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("Key '%s' was not found, using default value %t", key, defaultValue)
		return defaultValue
	}
	b, err := strconv.ParseBool(env)
	if err != nil {
		panic(err)
	}

	return b
}
