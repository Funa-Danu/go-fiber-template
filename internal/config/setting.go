package config

import (
	"os"
	"strconv"
)

// Setting holds environment based settings for go-fiber-template.
type Setting struct {
	ServiceName string
	Env         string
	Port        string

	ValkeyAddress string
	ValkeyDB      int

	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func mustGetEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", &MissingEnvError{Key: key}
	}
	return value, nil
}

// MissingEnvError indicates required env is not set.
type MissingEnvError struct {
	Key string
}

func (e *MissingEnvError) Error() string {
	return "missing required environment variable: " + e.Key
}

func mustAtoI(value string) (int, error) {
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return parsed, nil
}

func NewSetting() Setting {
	serviceName, _ := getEnvOrDefault("SERVICE_NAME", "go-fiber-template")
	env, _ := getEnvOrDefault("ENV", "local")
	port, _ := getEnvOrDefault("PORT", "3000")
	valkeyAddress, _ := getEnvOrDefault("VALKEY_ADDR", "localhost:6379")
	valkeyDB, _ := mustAtoI(getEnv("VALKEY_DB", "0"))
	pgHost, _ := getEnvOrDefault("PG_HOST", "localhost")
	pgPort, _ := getEnvOrDefault("PG_PORT", "5432")
	pgUser, _ := getEnvOrDefault("PG_USER", "postgres")
	pgPassword, _ := getEnvOrDefault("PG_PASSWORD", "postgres")
	pgDatabase, _ := getEnvOrDefault("PG_DATABASE", "postgres")

	return Setting{
		ServiceName:      serviceName,
		Env:              env,
		Port:             port,
		ValkeyAddress:    valkeyAddress,
		ValkeyDB:         valkeyDB,
		PostgresHost:     pgHost,
		PostgresPort:     pgPort,
		PostgresUser:     pgUser,
		PostgresPassword: pgPassword,
		PostgresDatabase: pgDatabase,
	}
}

func getEnvOrDefault(key, defaultValue string) (string, error) {
	if value := os.Getenv(key); value != "" {
		return value, nil
	}
	return defaultValue, nil
}

func (s Setting) validate() error {
	if s.ServiceName == "" {
		_, err := mustGetEnv("SERVICE_NAME")
		if err != nil {
			return err
		}
	}
	if s.Env == "" {
		_, err := mustGetEnv("ENV")
		if err != nil {
			return err
		}
	}
	if s.Port == "" {
		_, err := mustGetEnv("PORT")
		if err != nil {
			return err
		}
	}
	if s.ValkeyAddress == "" {
		_, err := mustGetEnv("VALKEY_ADDR")
		if err != nil {
			return err
		}
	}
	if _, err := strconv.Atoi(s.Port); err != nil {
		return err
	}
	if s.ValkeyDB < 0 {
		return &MissingEnvError{Key: "VALKEY_DB must be >= 0"}
	}
	if _, err := strconv.Atoi(s.PostgresPort); err != nil {
		return &MissingEnvError{Key: "PG_PORT must be int"}
	}
	if s.PostgresHost == "" {
		_, err := mustGetEnv("PG_HOST")
		if err != nil {
			return err
		}
	}
	return nil
}
