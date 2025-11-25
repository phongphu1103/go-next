package models

import "time"

type Article struct {
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
