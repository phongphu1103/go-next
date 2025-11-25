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

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, using system environment variables")
	}

	Cfg = Config{
		Connection: os.Getenv("DB_DRIVER"),
		Username:   os.Getenv("DB_USER"),
		Password:   os.Getenv("DB_PASS"),
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		Database:   os.Getenv("DB_NAME"),
		Charset:    os.Getenv("DB_CHARSET"),
		Collation:  os.Getenv("DB_COLLATION")
	}
}
