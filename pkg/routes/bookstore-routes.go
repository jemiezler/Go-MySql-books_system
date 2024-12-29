package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jemiezler/Go-MySql-books_system/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router fiber.Router) {
	router.Post("/books", controllers.CreateBook)
	router.Get("/books", controllers.FindAllBooks)
	router.Get("/books/:id", controllers.FindBookByID)
	router.Patch("/books/:id", controllers.UpdateBookByID)
	router.Delete("/books/:id", controllers.DeleteBookByID)
}
