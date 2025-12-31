package database

import "gorm.io/gorm"

type DatabaseManager struct {
    db *gorm.DB
}

var DBM *DatabaseManager

func NewDatabaseManager(db *gorm.DB) {
    DBM = &DatabaseManager{db: db}
}

/**
 * Get a database connection instance.
 */
func (m *DatabaseManager) Connection() *gorm.DB {
    return m.db
}

func (m *DatabaseManager) Disconnect() error {
    if m.db == nil {
        return nil
    }

    db, err := m.db.DB()
    if err != nil {
        return err
    }
    db.Close()
    // drop the connection instance
    m.db = nil

    return nil
}