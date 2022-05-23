package entities

import "time"

type Post struct {
	ID          int64     `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"title"`
	Content     string    `json:"content" gorm:"content"`
	PublishedAt time.Time `json:"published_at" gorm:"published_at"`
	CreatedAt   time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"updated_at"`
}
