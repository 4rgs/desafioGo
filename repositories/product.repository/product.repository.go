package product_repository

import (
	"context"
	"desafioGo/database"
	m "desafioGo/models"
	utils_service "desafioGo/services/utils.service"
	"math"
	"strconv"
	"sync"

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
	regex := "$regex"
	filter := bson.M{}
	findOptions := options.Find()

	if query := c.Query("query"); query != "" {
		filter = bson.M{
			"$or": []bson.M{
				{
					"brand": bson.M{
						regex: primitive.Regex{
							Pattern: query,
							Options: "i",
						},
					},
				},
				{
					"description": bson.M{
						regex: primitive.Regex{
							Pattern: query,
							Options: "i",
						},
					},
				},
				{
					"id": bson.M{
						regex: primitive.Regex{
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

	var wg sync.WaitGroup
	for cursor.Next(ctx) {
		wg.Add(1)
		cursor := cursor
		var product *m.Product
		var dsc m.Product
		go func() {
			defer wg.Done()
			cursor.Decode(&product)
			dsc = utils_service.ApplyDiscount(*product)
			prodWithDiscount = append(prodWithDiscount, dsc)
			products = append(products, product)

		}()
		wg.Wait()
	}

	return c.JSON(fiber.Map{
		"data":      products,
		"dsc":       prodWithDiscount,
		"total":     total,
		"page":      page,
		"last_page": math.Ceil(float64(total / perPage)),
	})
}
