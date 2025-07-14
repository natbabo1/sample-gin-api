package book

import "time"

type Book struct {
	ID          int64     `gorm:"primaryKey"`
	Title       string    `gorm:"size:200;not null"`
	Author      string    `gorm:"size:100;not null"`
	PublishedAt time.Time `gorm:"not null"`
	ISBN        string    `gorm:"size:20;not null;unique"`
	Pages       int       `gorm:"not null"`
	Language    string    `gorm:"size:50;not null"`
	Publisher   string    `gorm:"size:100;not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
