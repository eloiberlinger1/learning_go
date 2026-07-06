package env

import "os"

func GetString(key string, fallback string) string {
	if val, exists := os.LookupEnv(key); exists != false {
		return val
	}
	return fallback
}
