package config

import (
	"os"
	"strconv"
)

type DBConfig struct {
	DbName     string
	DbHost     string
	DbUser     string
	DbPassword string
	DbPort     string
	SSL        string
}

type Config struct {
	DBCfg     DBConfig
	Addr      string
	Timezone  string
	JWTSecret []byte
}

func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func GetEnvAsInt(name string, defaultVal int) int {
	valueStr := GetEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}
