package models

type User struct {
	ID       uint    `gorm:"primaryKey"`
	Email    string  `gorm:"size:100;not null;unique"`
	Password string  `gorm:"size:255;not null"`
}