package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Connection string
	Username   string
	Password   string
	Host       string
	Port       string
	Database   string
	Charset    string
	Collation  string
}

var Cfg Config

func GetEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, using system environment variables")
	}

	Cfg = Config{
		Connection: GetEnv("DB_CONNECTION", "mysql"),
		Username:   GetEnv("DB_USER", "forge"),
		Password:   GetEnv("DB_PASS", ""),
		Host:       GetEnv("DB_HOST", "localhost"),
		Port:       GetEnv("DB_PORT", "3306"),
		Database:   GetEnv("DB_NAME", "forge"),
		Charset:    GetEnv("DB_CHARSET", "utf8mb4"),
		Collation:  GetEnv("DB_COLLATION", "utf8mb4_unicode_ci"),
	}
}
