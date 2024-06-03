package env

import (
	"os"
	"strings"
)

func Load() error {
	content, err := os.ReadFile(".env")
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return err
		}

		key, value := parts[0], parts[1]
		err := os.Setenv(key, value)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		// return string and error instead of just string ?
		return ""
	}
	return value
}
