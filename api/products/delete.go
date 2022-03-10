package products

import "github.com/gofiber/fiber/v2"

func DeleteProducts(c *fiber.Ctx) error {
	return c.Status(200).SendString("Delete Product")
}
