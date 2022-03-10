package products

import "github.com/gofiber/fiber/v2"

func UpdateProducts(c *fiber.Ctx) error {
	return c.Status(200).SendString("Products")
}
