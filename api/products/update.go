package products

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ruiborda/go-mongodb-crud-api-rest/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func UpdateProducts(c *fiber.Ctx) error {
	_id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	var product Model
	err2 := c.BodyParser(&product)
	if err2 != nil {
		return c.Status(400).SendString(err2.Error())
	}

	m := database.Mongo{}
	m.Open()
	defer m.Close()

	result, err3 := m.Database.Collection("products").UpdateOne(
		c.Context(),
		bson.D{{"_id", _id}},
		bson.D{{"$set", bson.D{
			{"name", product.Name},
			{"description", product.Description},
			{"price", product.Price},
			{"updated_at", primitive.NewDateTimeFromTime(time.Now())},
		}}},
	)
	if err3 != nil {
		return c.Status(500).SendString(err3.Error())
	}

	if result.ModifiedCount == 0 {
		return c.Status(404).SendString("Product not found")
	}

	return c.Status(201).JSON(result)
}
