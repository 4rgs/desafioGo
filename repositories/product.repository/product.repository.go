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

func Find(c *fiber.Ctx) error {
	var products m.Products
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

	page, _ := strconv.Atoi(c.Query("page", "1"))
	var perPage int64 = 10

	total, _ := collection.CountDocuments(ctx, filter)

	findOptions.SetSkip((int64(page) - 1) * perPage)
	findOptions.SetLimit(perPage)

	cursor, _ := collection.Find(ctx, filter, findOptions)
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product *m.Product
		cursor.Decode(&product)
		products = append(products, product)
	}

	return c.JSON(fiber.Map{
		"data":      products,
		"total":     total,
		"page":      page,
		"last_page": math.Ceil(float64(total / perPage)),
	})
}
