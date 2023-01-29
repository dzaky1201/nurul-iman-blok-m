package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"nurul-iman-blok-m/model"
	"os"
)

func Db() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal(errEnv.Error())
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	errMigrate := db.AutoMigrate(&model.User{}, &model.Role{}, &model.Announcement{}, &model.Article{}, &model.Category{}, &model.StudyRundown{}, &model.StudyVideo{})
	if errMigrate != nil {
		log.Fatal(errMigrate.Error())
	}

	fmt.Println("database connected")

	return db
}
