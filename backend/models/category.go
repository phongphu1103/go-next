package models

import (
    // "gorm.io/gorm"
    "time"
)

type Category struct {
    // gorm.Model // automatically add common fields
    ID        uint      `gorm:"primaryKey"`
    Name      string    `gorm:"size:100;not null"`
    Slug      string    `gorm:"size:100;not null;unique"`
    Articles  []Article `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
