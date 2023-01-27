package model

import (
	"book_store_api/database"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `gorm:"size:255;not null;unique" json:"title"`
	Author string `gorm:"size:255;not null" json:"author"`
	Genre  string `gorm:"size:100;not null" json:"genre"`
	UserID uint
}

func (book *Book) Save() (*Book, error) {
	err := database.Database.Create(&book).Error
	if err != nil {
		return &Book{}, err
	}
	return book, nil
}
