package main

import (
	product_service "desafioGo/services/product.service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hola Mundo, by 4rGs!")
	})
	app.Get("/api/productos/busqueda", func(c *fiber.Ctx) error {
		return product_service.Find(c)
	})
	app.Listen(":8080")
}
