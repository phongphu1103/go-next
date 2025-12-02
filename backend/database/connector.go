package database

import (
    "fmt"
    "log"
    "time"
    "sync"

    "gorm.io/gorm"
    "backend/config"
)

// use a type alias
type Config = config.Config

var (
    db   *gorm.DB
    once sync.Once
)

type Driver interface {
    GetDSN(cfg Config) string
    Open(dsn string) (*gorm.DB, error)
}

var driverRegistry = map[string]Driver{
    "mysql":    MySQLDriver{},
    "postgres": PostgresDriver{},
}

/**
 * Create a connector instance based on the configuration.
 */
func createConnector(name string) (Driver, error) {
    // If found: ok = true, and d = driver instance
    // If not found: ok = false, and d = zero-value
    if d, ok := driverRegistry[name]; ok {
        return d, nil
    }

    return nil, fmt.Errorf("Unsupported driver: %s", name)
}

func Connect(cfg Config) (*gorm.DB, error) {
    var err error
    driver, err := createConnector(cfg.Connection)
    if err != nil {
        return nil, err
    }

    once.Do(func() {
        dsn := driver.GetDSN(cfg)
        db, err = driver.Open(dsn)
        if err != nil {
            return
        }

        sqlDB, _ := db.DB()
        sqlDB.SetMaxIdleConns(10)
        sqlDB.SetMaxOpenConns(100)
        sqlDB.SetConnMaxLifetime(time.Hour)
        log.Printf("Connected to %T", driver)
    })

    // after once.Do() finishes
    return db, err
}
