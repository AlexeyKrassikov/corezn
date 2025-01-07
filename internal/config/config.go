package config

import "os"

type Config struct {
        ServerPort   string
        Environment  string
        DatabaseURL  string
        DatabaseName string
}

func LoadConfig() *Config {
        return &Config{
                ServerPort:   getEnv("SERVER_PORT", "8080"),
                Environment:  getEnv("ENVIRONMENT", "development"),
                DatabaseURL:  getEnv("DATABASE_URL", "postgresql://postgres:postgres@localhost:5432/corezn?sslmode=disable"),
                DatabaseName: getEnv("DATABASE_NAME", "corezn"),
        }
}

func getEnv(key, defaultValue string) string {
        if value, exists := os.LookupEnv(key); exists {
                return value
        }
        return defaultValue
}