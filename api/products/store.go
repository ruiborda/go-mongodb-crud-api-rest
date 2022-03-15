package products

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ruiborda/go-mongodb-crud-api-rest/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func StoreProducts(c *fiber.Ctx) error {
	var product Model
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	//force generate new id
	product.ID = primitive.NewObjectID()

	m := database.Mongo{}
	m.Open()
	defer m.Close()

	result, err2 := m.Database.Collection("products").InsertOne(c.Context(), product)
	if err2 != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(201).JSON(result)
}
