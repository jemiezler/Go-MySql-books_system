package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jemiezler/Go-MySql-books_system/pkg/routes"
)

func main() {
	app := fiber.New()
	api := app.Group("/api")
	routes.RegisterBookStoreRoutes(api)
	log.Fatal(app.Listen(":8080"))
}
