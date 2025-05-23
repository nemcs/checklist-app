package config

import "os"

const (
	DBHost        = "DBHost"
	DBPort        = "DBPort"
	DBUser        = "DBUser"
	DBPass        = "DBPass"
	DBName        = "DBName"
	DefaultEnvVal = ""
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
			DBHost: getEnv(DBHost, DefaultEnvVal),
			DBPort: getEnv(DBPort, DefaultEnvVal),
			DBUser: getEnv(DBUser, DefaultEnvVal),
			DBPass: getEnv(DBPass, DefaultEnvVal),
			DBName: getEnv(DBName, DefaultEnvVal),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
