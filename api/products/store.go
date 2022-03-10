package products

import (
	"github.com/gofiber/fiber/v2"
)

func StoreProducts(c *fiber.Ctx) error {
	/*	product := new(Model)
		if err := c.BodyParser(product); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		/*	product.ID = ""

			// insert the record
			insertionResult, err := collection.InsertOne(c.Context(), product)
			if err != nil {
				return c.Status(500).SendString(err.Error())
			}

			// get the just inserted record in order to return it as response
			filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
			createdRecord := collection.FindOne(c.Context(), filter)

			// decode the Mongo record into Employee
			createdEmployee := &Model{}
			createdRecord.Decode(createdEmployee)
	*/
	// return the created Employee in JSON format*/
	return c.Status(201).JSON(c.Body())
}
