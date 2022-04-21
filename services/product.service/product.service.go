package product_service

import (
	m "desafioGo/models"
	prodRepository "desafioGo/repositories/product.repository"
	"fmt"
)

func Read() (m.Products, error) {

	products, err := prodRepository.ListAll()

	if err != nil {
		return nil, err
	}

	return products, nil
}
func Find(query string) (m.Products, error) {

	products, err := prodRepository.Find(query)
	fmt.Println(products)
	if err != nil {
		return nil, err
	}

	return products, nil
}
func applyDiscount(product m.Product) m.Product {
	if productIsPalidrom(product) {
		product.Price = product.Price * 0.5
	}
	return product
}

func productIsPalidrom(product m.Product) bool {
	return (isPalindrom(string(product.ID)) || isPalindrom(product.Brand) || isPalindrom(product.Description))
}

func isPalindrom(query string) bool {
	wordLen1 := len(query)
	for i := 0; i < wordLen1; i++ {
		if query[i] != query[wordLen1-1-i] {
			return false
		}
	}
	return true
}
