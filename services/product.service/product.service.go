package product_service

import (
	prodRepository "desafioGo/repositories/product.repository"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Find(c *fiber.Ctx) error {
	fmt.Println(prodRepository.Find(c), c)
	return nil
}
