package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ErrEnvNotFound is returned when a required environment variable is unset or empty.
type ErrEnvNotFound struct {
	Key string // Name of the environment variable
}

func (e ErrEnvNotFound) Error() string {
	return fmt.Sprintf("config: environment variable %q is required", e.Key)
}

// ErrInvalidEnv is returned when an environment variable is set to a value
// that cannot be parsed as the expected type.
type ErrInvalidEnv struct {
	Key   string // Name of the envionement variable
	Value string // Raw value that could not be parsed
	Err   error  // Underlying parse error
}

func (e ErrInvalidEnv) Error() string {
	return fmt.Sprintf(
		"config: environment variable %q has invalid value %q: %v",
		e.Key,
		e.Value,
		e.Err,
	)
}

// Unwrap returns the underlying parse error.
func (e ErrInvalidEnv) Unwrap() error { return e.Err }

// envValue returns the raw value of the environment variable named key and
// reports whether it was set to a non-blank value.
func envValue(key string) (string, bool) {
	value := os.Getenv(key)
	if strings.TrimSpace(value) == "" {
		return value, false
	}
	return value, true
}

// readRequiredString returns the value of the environment variable named key.
//
// It returns a MissingEnvError if the variable is unset or contains only
// whitespace.
func readRequiredString(key string) (string, error) {
	value, ok := envValue(key)
	if !ok {
		return "", ErrEnvNotFound{Key: key}
	}
	return value, nil
}

// readOptionalString returns the value of the environment variable named key,
// or fallback if the variable is unset or contains only whitespace.
func readOptionalString(key, fallback string) string {
	value, ok := envValue(key)
	if !ok {
		return fallback
	}
	return value
}

// readRequiredInt returns the value of the environment variable named key
// parsed as an int.
//
// It returns a MissingEnvError if the variable is unset or contains only
// whitespace, and an InvalidEnvError if it does not hold a valid integer.
func readRequiredInt(key string) (int, error) {
	value, ok := envValue(key)
	if !ok {
		return 0, ErrEnvNotFound{Key: key}
	}
	n, err := strconv.Atoi(strings.TrimSpace(value))
	if err != nil {
		return 0, ErrInvalidEnv{Key: key, Value: value, Err: err}
	}
	return n, nil
}

// readOptionalInt returns the value of the environment variable named key
// parsed as an int, or fallback if the variable is unset or contains only
// whitespace.
//
// It returns an InvalidEnvError if the variable is set but does not hold a
// valid integer.
func readOptionalInt(key string, fallback int) (int, error) {
	value, ok := envValue(key)
	if !ok {
		return fallback, nil
	}
	n, err := strconv.Atoi(strings.TrimSpace(value))
	if err != nil {
		return 0, ErrInvalidEnv{Key: key, Value: value, Err: err}
	}
	return n, nil
}
