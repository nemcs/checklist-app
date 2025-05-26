package config

import "os"

const (
	DBHost = "POSTGRES_HOST"
	DBPort = "POSTGRES_PORT"
	DBUser = "POSTGRES_USER"
	DBPass = "POSTGRES_PASSWORD"
	DBName = "POSTGRES_DB"
)

type PostgresConfig struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

type Config struct {
	PostgresConfig
}

func New() *Config {
	return &Config{
		PostgresConfig{
			DBHost: getEnv(DBHost, "postgresql"),
			DBPort: getEnv(DBPort, "5432"),
			DBUser: getEnv(DBUser, "postgres"),
			DBPass: getEnv(DBPass, "postgres"),
			DBName: getEnv(DBName, "checklist_app"),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
