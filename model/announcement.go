package model

import "time"

type Announcement struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Title       string `gorm:"size:255;not null"`
	Description string `gorm:"type:text;not null"`
	Images      string `gorm:"size:100;not null"`
	User        User
	UserID      uint   `gorm:"index;not null"`
	Slug        string `gorm:"size:255;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
