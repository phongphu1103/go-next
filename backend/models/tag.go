package models

import (
    "time"
)

type Tag struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `gorm:"size:50;unique"`
    Slug      string    `gorm:"size:50;not null"`
    CreatedAt time.Time
}

type ArticleTag struct {
    ArticleID uint `gorm:"primaryKey"`
    TagID     uint `gorm:"primaryKey"`
}
