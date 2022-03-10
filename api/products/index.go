package products

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ruiborda/shop/database"
	"go.mongodb.org/mongo-driver/bson"
)

func IndexProducts(c *fiber.Ctx) error {

	m := database.Mongo{}
	m.Open()
	defer m.Close()
	m.SetCollection("products")

	query := bson.D{{"title", "Back to the Future"}}

	result := m.FindOne(&query)

	fmt.Printf("%s\n", result)

	jsonData, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
	return c.Status(200).SendString("Products")
}
