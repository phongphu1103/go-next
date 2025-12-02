package models

import (
    "gorm.io/gorm"
    "time"
)

type Article struct {
    gorm.Model
    ID          uint      `gorm:"primaryKey"`
    Title       string    `gorm:"size:255;not null"`
    Slug        string    `gorm:"size:255;not null;unique"`
	Description string    `gorm:"size:255;not null"`
    Content     string    `gorm:"type:text;not null"`
    CategoryID  uint
    Category    Category
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
