package main

import (
	product_repository "desafioGo/repositories/product.repository"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/api/productos/busqueda", func(c *fiber.Ctx) error {
		return product_repository.Find(c)
	})
	app.Listen(":8080")
}
