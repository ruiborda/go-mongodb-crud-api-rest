package products

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/ruiborda/go-mongodb-crud-api-rest/database"
	"go.mongodb.org/mongo-driver/bson"
)

func IndexProducts(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json; charset=utf-8")
	m := database.Mongo{}
	m.Open()
	defer m.Close()

	cursor, err := m.Database.Collection("products").Find(context.TODO(), bson.D{{}})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var results []Model
	err1 := cursor.All(context.TODO(), &results)
	if err1 != nil {
		return c.Status(500).SendString(err1.Error())
	}

	marshal, err2 := json.Marshal(results)
	if err2 != nil {
		return c.Status(500).SendString(err2.Error())
	}

	return c.Status(200).SendString(string(marshal))
}
