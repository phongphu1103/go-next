package models

import "time"

type Category struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `gorm:"size:100;not null;unique"`
    Slug      string    `gorm:"size:100;not null;unique"`
    Articles  []Article `gorm:"foreignKey:CategoryID"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
