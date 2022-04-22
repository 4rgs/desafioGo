package product_repository

import (
	"context"
	"desafioGo/database"
	m "desafioGo/models"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection = database.GetCollection("products")
var ctx = context.Background()

//MODULARIZAR funcionalidades a service
func Find(c *fiber.Ctx) error {
	var products m.Products
	var prodWithDiscount m.DiscountProducts
	filter := bson.M{}
	findOptions := options.Find()

	if query := c.Query("query"); query != "" {
		filter = bson.M{
			"$or": []bson.M{
				{
					"brand": bson.M{
						"$regex": primitive.Regex{
							Pattern: query,
							Options: "i",
						},
					},
				},
				{
					"description": bson.M{
						"$regex": primitive.Regex{
							Pattern: query,
							Options: "i",
						},
					},
				},
				{
					"id": bson.M{
						"$regex": primitive.Regex{
							Pattern: query,
							Options: "i",
						},
					},
				},
			},
		}
	}

	if sort := c.Query("sort"); sort != "" {
		if sort == "asc" {
			findOptions.SetSort(bson.D{{"price", 1}})
		} else if sort == "desc" {
			findOptions.SetSort(bson.D{{"price", -1}})
		}
	}

	page, _ := strconv.Atoi(c.Query("page", "0"))
	var perPage int64 = 10
	total, _ := collection.CountDocuments(ctx, filter)

	findOptions.SetSkip((int64(page)) * perPage)
	findOptions.SetLimit(perPage)

	cursor, _ := collection.Find(ctx, filter, findOptions)
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product *m.Product
		var dsc m.Product
		cursor.Decode(&product)
		dsc = applyDiscount(*product)
		prodWithDiscount = append(prodWithDiscount, dsc)
		products = append(products, product)
	}
	return c.JSON(fiber.Map{
		"data":      products,
		"dsc":       prodWithDiscount,
		"total":     total,
		"page":      page,
		"last_page": math.Ceil(float64(total / perPage)),
	})
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
