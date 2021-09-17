package utils

import "os"

func GetEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	} else {
		err := os.Setenv(key, defaultValue)
		if err != nil {
			return defaultValue
		}
		return os.Getenv(key)
	}
}
