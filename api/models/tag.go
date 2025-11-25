type Tag struct {
    ID   uint   `gorm:"primaryKey"`
    Name string `gorm:"size:50;unique"`
}

type ArticleTag struct {
    ArticleID uint
    TagID     uint
}
