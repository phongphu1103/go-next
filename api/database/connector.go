package database

import (
	"errors"
	"api/config"

	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

type Driver interface {
	GetDSN(cfg Config) string
	CreateConnection(dsn string) (*gorm.DB, error)
}

func Connect() (*gorm.DB, error) {
	switch config.Cfg.Connection {
		case "mysql":
			
		case "postgres":
		
	}
}