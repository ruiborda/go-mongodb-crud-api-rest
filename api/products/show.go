package products

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ruiborda/go-mongodb-crud-api-rest/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ShowProducts(c *fiber.Ctx) error {
	_id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	m := database.Mongo{}
	m.Open()
	defer m.Close()

	var result Model

	if err := m.Database.Collection("products").FindOne(c.Context(), bson.M{"_id": _id}).Decode(&result); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
