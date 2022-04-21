package product_service

import (
	prodRepository "desafioGo/repositories/product.repository"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Find(c *fiber.Ctx) error {
	products := prodRepository.Find(c)

	fmt.Println(products, &products)
	return products
}
