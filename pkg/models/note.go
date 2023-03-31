package models

import (
	"time"

	"github.com/hammad-umar/goland-gin-crud-api/pkg/config"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	ID 			uint		`gorm:"primaryKey" json:"id"`
	Title 		string 		`json:"title"`	
	Description string 		`json:"description"`
	CreatedAt   time.Time	`json:"created_at"`
    UpdatedAt   time.Time	`json:"updated_at"`
}

var db *gorm.DB 

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Note{})
}

func Find() []Note {
	var notes []Note
	db.Find(&notes)
	return notes 
}

func FindOneById(id int64) *Note {
	var note Note
	db.Where("id = ?", id).Find(&note)
	return &note 
}

func (n *Note) Create() *Note {
	db.Create(n)
	return n 
}

func DeleteOneById(id int64) *Note {
	var note Note 
	db.Where("id = ?", id).Delete(&note)
	return &note 
}
