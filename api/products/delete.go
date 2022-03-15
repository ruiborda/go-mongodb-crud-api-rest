package products

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ruiborda/go-mongodb-crud-api-rest/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteProducts(c *fiber.Ctx) error {
	_id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	m := database.Mongo{}
	m.Open()
	defer m.Close()

	one, err2 := m.Database.Collection("products").DeleteOne(c.Context(), bson.M{"_id": _id})
	if err2 != nil {
		return c.Status(500).SendString(err.Error())
	}
	if one.DeletedCount == 0 {
		return c.Status(404).SendString("Product not found")
	}

	return c.Status(200).SendString("Delete Product")
}
