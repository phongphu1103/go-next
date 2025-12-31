package models

import (
    // "gorm.io/gorm"
    "time"
)

type Article struct {
    // gorm.Model
    ID          uint      `gorm:"primaryKey"`
    Title       string    `gorm:"size:255;not null"`
    Slug        string    `gorm:"size:255;not null;uniqueIndex"`
	Description string    `gorm:"size:255;not null"`
    Content     string    `gorm:"type:text;not null"`
    CategoryID  uint      `gorm:"not null"`
    Category    *Category  `gorm:"foreignKey:CategoryID"`
    Tags        []*Tag     `gorm:"many2many:article_tags"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
