package model

import "time"

type Role struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	RoleName  string `gorm:"type:varchar(100);NOT NULL"`
	Users     []User
	CreatedAt time.Time
	UpdatedAt time.Time
}
