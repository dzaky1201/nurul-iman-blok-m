package main

import (
	"nurul-iman-blok-m/database"
	"nurul-iman-blok-m/model"
)

func DatabaseMigration() {
	db := database.Db()
	errMigrate := db.AutoMigrate(&model.User{}, &model.Role{}, &model.Announcement{}, &model.Article{}, &model.Category{}, &model.StudyRundown{}, &model.StudyVideo{})
	if errMigrate != nil {
		return
	}
}

func main() {
	DatabaseMigration()
}
