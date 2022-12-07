package model

import "time"

type StudyVideo struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;not null"`
	Title     string `gorm:"size:100;not null"`
	thumbnail string `gorm:"size:100;not null"`
	Url       string `gorm:"size:255;not null"`
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
