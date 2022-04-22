package product_service

import (
	prodRepository "desafioGo/repositories/product.repository"

	"github.com/gofiber/fiber/v2"
)

func Find(c *fiber.Ctx) error {
	products := prodRepository.Find(c)
	return products
}
