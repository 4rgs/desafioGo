package product_service

import m "desafioGo/models"

func ProductIsPalidrom(product m.Product) bool {
	return (IsPalindrom(string(product.ID)) || IsPalindrom(product.Brand) || IsPalindrom(product.Description))
}
