package database

import "gorm.io/gorm"

// This exposes the DB to the rest of the application
func DB() *gorm.DB {
    return DBM.Connection()   // DBM is your DatabaseManager singleton
}

func Disconnect() error {
    return DBM.Disconnect()
}