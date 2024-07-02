package env

import (
	"os"
	"strings"
)

func Load() {
	// Check if .env file exists
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		return // .env file does not exist, skip loading
	}

	// Load .env file content
	content, err := os.ReadFile(".env")
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key, value := parts[0], parts[1]
		os.Setenv(key, value)
	}
}

func GetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		// return string and error instead of just string ?
		return ""
	}
	return value
}
