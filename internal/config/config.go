package config

import (
	"os"
	"strconv"
)

type DBConfig struct {
	DbURL  string
	DbType int
}

type Config struct {
	DBCfg DBConfig
}

func NewConfig() *Config {
	return &Config{
		DBCfg: DBConfig{
			DbURL:  getEnv("DB_URL", ""),
			DbType: getEnvAsInt("DB_TYPE", 0),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}
