package utiles

import "os"

var Environment string

func init() {
	Environment = GetEnv("EVO_ENV", "dev")
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}

	return value
}
