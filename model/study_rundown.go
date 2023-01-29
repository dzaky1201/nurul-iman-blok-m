package model

import "time"

type StudyRundown struct {
	ID           uint   `gorm:"primaryKey;autoIncrement;not null"`
	Title        string `gorm:"size:100;not null"`
	OnScheduled  bool   `gorm:"type:boolean"`
	ScheduleDate string `gorm:"size:100; not null"`
	User         User
	UserID       uint   `gorm:"index;not null"`
	Time         string `gorm:"size:100;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
