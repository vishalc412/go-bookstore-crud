package models

import (
	"github.com/jinzhu/gorm"
	"github.com/vishalc412/go-bookstore-crud/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"name" json:"name"`
	Author      string `gorm:"author" json:"author"`
	Publication string `gorm:"publication" json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	books := []Book{
		{Name: "Book 1", Author: "Author 1", Publication: "Publication 1"},
		{Name: "Book 2", Author: "Author 2", Publication: "Publication 2"},
		{Name: "Book 3", Author: "Author 3", Publication: "Publication 3"},
	}
	db.AutoMigrate(books)
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBook() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("Id=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int) Book {
	var book Book
	db.Where("Id=?", Id).Delete(&book)
	return book
}
