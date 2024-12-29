package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jemiezler/Go-MySql-books_system/pkg/models"
)

var NewBook models.Book

func CreateBook(c *fiber.Ctx) error {
	var book models.Book

	// Parse the request body into the Book struct
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to parse request body",
		})
	}

	//Create the book in the database
	createdBook := book.CreateBook()
	return c.Status(fiber.StatusCreated).JSON(createdBook)
}

func FindAllBooks(c *fiber.Ctx) error {
	Books := models.FindAllBooks()
	return c.Status(fiber.StatusOK).JSON(Books)
}

func FindBookByID(c *fiber.Ctx) error {
	id := c.Params("id")
	bookID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}
	book := models.FindBookByID(bookID)
	if book == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Book not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(book)
}

func UpdateBookByID(c *fiber.Ctx) error {
	id := c.Params("id")
	bookID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}

	var updatedBook models.Book
	if err := c.BodyParser(&updatedBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to parse request body",
		})
	}

	book := models.UpdateBookByID(bookID, updatedBook)
	return c.Status(fiber.StatusOK).JSON(book)
}

func DeleteBookByID(c *fiber.Ctx) error {
	id := c.Params("id")
	bookID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}

	book := models.DeleteBookByID(bookID)
	if book.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Book deleted successfully",
		"id":      book.ID,
	})
}
