package osenv

import (
	"os"
	"strings"
)

// Get environment variable, or the supplied alternative if unset.
func Get(key, alt string) string {
	v := os.Getenv(key)
	if v == "" {
		return alt
	}

	return v
}

// GetArray of strings from one environment variable.
// Returns an empty array if the envvar is unset.
func GetArray(key string) []string {
	v := os.Getenv(key)
	if v == "" {
		return []string{}
	}

	return strings.Split(v, ",")
}
