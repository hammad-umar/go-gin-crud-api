package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID 			uint		`gorm:"primaryKey" json:"id"`
	Name 		string 		`json:"name"`	
	Email		string 		`gorm:"unique" json:"email"`
	Password	string 		`json:"password"`
	CreatedAt   time.Time	`json:"created_at"`
    UpdatedAt   time.Time	`json:"updated_at"`
}

func init() {
	db.AutoMigrate(&User{})
}
