package models

type Tag struct {
    ID   uint   `gorm:"primaryKey"`
    Name string `gorm:"size:50;unique"`
    Slug string `gorm:"size:255;not null`
}

type ArticleTag struct {
    ArticleID uint
    TagID     uint
}
