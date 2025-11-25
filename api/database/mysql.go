package database

import (
	"fmt"
	"log"
	"time"
	
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"api/config"
)

var (
	db   *gorm.DB
	once sync.Once
)

/**
 * Create a DSN string from a configuration.
 */
func GetDSN(cfg Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.Charset)
}

/**
 * Create a new connection.
 */
func CreateConnection(dsn string) (*gorm.DB, error) {
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Set connection pool values
	connector, _ := connection.DB()
	connector.SetMaxIdleConns(10)
	connector.SetMaxOpenConns(100)
	connector.SetConnMaxLifetime(time.Hour)

	return connection, nil
}

func Connect() (*gorm.DB, error) {
	var err error
	cfg := config.Cfg

	once.Do(func() {
		dsn := GetDSN(cfg)
		var connection *gorm.DB
		connection, err = CreateConnection(dsn)
		if err != nil {
			return
		}

		db = connection
		log.Println("Connected to MySQL")
	})

	// after once.Do() finishes
	return db, err
}