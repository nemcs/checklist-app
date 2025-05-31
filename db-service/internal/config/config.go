package config

import "os"

const (
	Host   = "POSTGRES_HOST"
	Port   = "POSTGRES_PORT"
	User   = "POSTGRES_USER"
	Pass   = "POSTGRES_PASSWORD"
	DBName = "POSTGRES_DB"
)

type PostgresConfig struct {
	Host   string
	Port   string
	User   string
	Pass   string
	DBName string
}

type Config struct {
	PostgresConfig
}

func New() *Config {
	return &Config{
		PostgresConfig{
			Host:   getEnv(Host, "postgresql"),
			Port:   getEnv(Port, "5432"),
			User:   getEnv(User, "postgres"),
			Pass:   getEnv(Pass, "postgres"),
			DBName: getEnv(DBName, "checklist_app"),
		},
	}
}

func getEnv(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
