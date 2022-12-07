package model

import "time"

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"type:varchar(100);NOT NULL"`
	Email    string `gorm:"type:varchar(100);NOT NULL"`
	Password string `gorm:"type:varchar(255);NOT NULL"`
	// for migration
	Role          Role
	RoleID        uint `gorm:"index;NOT NULL"`
	Announcements []Announcement
	Articles      []Article
	StudyRundown  []StudyRundown
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
