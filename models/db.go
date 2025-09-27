package models

import (
	"log"
	"web3-go-blog/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open(config.DBPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	err = DB.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}

func ClearDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open(config.DBPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	DB.Exec("DROP TABLE users")
	DB.Exec("DROP TABLE posts")
	DB.Exec("DROP TABLE comments")
}
