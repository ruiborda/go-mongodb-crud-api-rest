package products

import "github.com/gofiber/fiber/v2"

func StoreProducts(c *fiber.Ctx) error {
	return c.Status(200).SendString("Create Product")
}
