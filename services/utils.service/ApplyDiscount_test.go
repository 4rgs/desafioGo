package product_service

import (
	m "desafioGo/models"
	"reflect"
	"testing"
)

var palindromo, noPalindromo, palindromoMod m.Product

func iniciarVariables() {
	//Producto Original con Brand Palindrome
	palindromo.ID = 154
	palindromo.Brand = "assa"
	palindromo.Description = "ahelf lxhñep"
	palindromo.Image = "www.lider.cl/catalogo/images/tvIcon.svg"
	palindromo.Price = 748655
	//Producto con descunto aplicado
	palindromoMod.ID = 154
	palindromoMod.Brand = "assa"
	palindromoMod.Description = "ahelf lxhñep"
	palindromoMod.Image = "www.lider.cl/catalogo/images/tvIcon.svg"
	palindromoMod.Price = 374327.5
	//Producto  no palindrome
	noPalindromo.ID = 155
	noPalindromo.Brand = "lcjuwnr"
	noPalindromo.Description = "wrrgd jpzmzl"
	noPalindromo.Image = "www.lider.cl/catalogo/images/toysIcon.svg"
	noPalindromo.Price = 748655
}

func TestApplyDiscount(t *testing.T) {
	iniciarVariables()
	type args struct {
		product m.Product
	}
	tests := []struct {
		name string
		args args
		want m.Product
	}{
		{
			name: "Testing with Prod 154 must return product with discount",
			args: args{palindromo},
			want: palindromoMod,
		},
		{
			name: "Testing with Prod 155 must return the same product",
			args: args{noPalindromo},
			want: noPalindromo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ApplyDiscount(tt.args.product); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApplyDiscount() = %v, want %v", got, tt.want)
			}
		})
	}
}
