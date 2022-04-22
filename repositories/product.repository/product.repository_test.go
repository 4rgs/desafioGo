package product_repository

import (
	m "desafioGo/models"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var testProducts m.Products

var testDsc m.DiscountProducts

func TestFind(t *testing.T) {
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		// First test case
		{
			description:  "get HTTP status 200",
			route:        "/api/productos/busqueda",
			expectedCode: 200,
		},
		// Second test case
		{
			description:  "get HTTP status 404, when route is not exists",
			route:        "/not-found",
			expectedCode: 404,
		},
	}
	// Define Fiber app.
	app := fiber.New()

	// Create route with GET method for test
	app.Get("/api/productos/busqueda", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data":      testProducts,
			"dsc":       testDsc,
			"total":     600,
			"page":      1,
			"last_page": 6000,
		})
	})

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route from the test case
		req := httptest.NewRequest("GET", test.route, nil)

		// Perform the request plain with the app,
		// the second argument is a request latency
		// (set to -1 for no latency)
		resp, _ := app.Test(req, 1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
