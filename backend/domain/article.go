package domain

import "time"

type Article struct {
    ID          uint
    Title       string
    Slug        string
    Description string
    Content     string
    CategoryID  uint
    Category    *Category
    Tags        []*Tag
    CreatedAt   time.Time
}