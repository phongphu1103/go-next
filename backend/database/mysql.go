package database

import (
	"fmt"
	
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLDriver struct{}

/**
 * Create a DSN string from a configuration.
 */
func (MySQLDriver) GetDSN(cfg Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.Charset)
}

/**
 * Create a new connection.
 */
func (MySQLDriver) Open(dsn string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
