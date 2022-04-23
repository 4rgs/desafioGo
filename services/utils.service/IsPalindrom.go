package product_service

func IsPalindrom(query string) bool {
	wordLen1 := len(query)
	for i := 0; i < wordLen1; i++ {
		if query[i] != query[wordLen1-1-i] {
			return false
		}
	}
	return true
}
