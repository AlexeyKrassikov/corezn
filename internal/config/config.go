package config

import "os"

type Config struct {
        ServerPort string
        Environment string
}

func LoadConfig() *Config {
        return &Config{
                ServerPort: getEnv("SERVER_PORT", "8080"),
                Environment: getEnv("ENVIRONMENT", "development"),
        }
}

func getEnv(key, defaultValue string) string {
        if value, exists := os.LookupEnv(key); exists {
                return value
        }
        return defaultValue
}