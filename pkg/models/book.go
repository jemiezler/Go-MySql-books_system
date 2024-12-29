package models

import (
	"github.com/google/uuid"
	"github.com/jemiezler/Go-MySql-books_system/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Name        string    `json:"name"`
	Author      string    `json:"author"`
	Publication string    `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	err := db.AutoMigrate(&Book{})
	if err != nil {
		panic("Failed to migrate the database schema: " + err.Error())
	}
}

func (b *Book) CreateBook() *Book {
	b.ID = uuid.New()
	db.Create(b)
	return b
}

func FindAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func FindBookByID(id uuid.UUID) *Book {
	var book Book
	db.First(&book, "id = ?", id)
	return &book
}

func DeleteBookByID(id uuid.UUID) Book {
	var book Book
	db.First(&book, "id = ?", id).Delete(book)
	return book
}

func UpdateBookByID(id uuid.UUID, updatedBook Book) *Book {
	var book Book
	// Find the existing book by UUID
	db.First(&book, "id = ?", id)
	// Update the fields of the book
	book.Name = updatedBook.Name
	book.Author = updatedBook.Author
	book.Publication = updatedBook.Publication
	// Save the updated book
	db.Save(&book)
	return &book
}
