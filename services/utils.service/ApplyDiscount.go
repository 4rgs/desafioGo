package product_service

import m "desafioGo/models"

func ApplyDiscount(product m.Product) m.Product {
	if ProductIsPalidrom(product) {
		product.Price = product.Price * 0.5
	}
	return product
}
