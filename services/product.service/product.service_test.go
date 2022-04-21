package product_service_test

import (
	product_service "desafioGo/services/product.service"
	"testing"
)

func TestFind(t *testing.T) {
	productos, err := product_service.Find("dsaasd")

	if err != nil {
		t.Error("/productos/busqueda Find()	EXPECT []Products, GOT no response from server")
		t.Fail()
	}

	if len(productos) == 0 {
		t.Error("/productos/busqueda Find()	EXPECT []Products, GOT 	empty array")
		t.Fail()
	} else {
		t.Log("/productos/busqueda Find()	EXPECT []Products, GOT []Products")
	}
}
func TestGetAll(t *testing.T) {
	productos, err := product_service.Find()

	if err != nil {
		t.Error("/productos/busqueda Find()	EXPECT []Products, GOT no response from server")
		t.Fail()
	}

	if len(productos) == 0 {
		t.Error("/productos/busqueda Find()	EXPECT []Products, GOT 	empty array")
		t.Fail()
	} else {
		t.Log("/productos/busqueda Find()	EXPECT []Products, GOT []Products")
	}
}
