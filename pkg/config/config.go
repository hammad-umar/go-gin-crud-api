package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=78678612 dbname=gincrud port=5432 sslmode=disable TimeZone=Asia/Karachi"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db 
}