package env

import (
	"fmt"
	"os"
)

const (
	EnvPort        = "PORT"
	EnvLoggerLevel = "LOGGER_LEVEL"
	EnvOrigin      = "ORIGIN"

	DefaultPort        = "8080"
	DefaultLoggerLevel = "DEBUG"
	DefaultOrigin      = "*"
)

func GetEnv(args ...string) string {
	value := os.Getenv(args[0])
	if value == "" && len(args) > 1 {
		return args[1]
	}

	return value
}

func CheckRequiredEnv(envs ...string) error {
	for _, env := range envs {
		value := os.Getenv(env)
		if value == "" {
			return &EnvRequiredError{
				Env: env,
			}
		}
	}

	return nil
}

type EnvRequiredError struct {
	Env string
}

func (err *EnvRequiredError) Error() string {
	return fmt.Sprintf("env %s is required", err.Env)
}
