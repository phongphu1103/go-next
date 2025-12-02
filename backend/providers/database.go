package providers

import (
    "fmt"
    "gorm.io/gorm"

    "backend/config"
    "backend/database"
)

type DatabaseServiceProvider struct{
    DB *gorm.DB
}

func (DatabaseServiceProvider) Register() {
    // Load environment variables
    config.LoadConfig()
}

func (p *DatabaseServiceProvider) Boot() error {
    // Connect to database
    db, err := database.Connect(config.Cfg)
    if err != nil {
        return fmt.Errorf("Database connection failed: %w", err)
    }
    p.DB = db
    return nil
}
