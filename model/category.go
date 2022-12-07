package model

import "time"

type Category struct {
	ID           uint   `gorm:"autoIncrement;not null;primaryKey"`
	CategoryName string `gorm:"size(100);not null"`
	Articles     []Article
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
