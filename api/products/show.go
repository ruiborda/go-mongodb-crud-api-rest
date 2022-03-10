package products

import "github.com/gofiber/fiber/v2"

func ShowProducts(c *fiber.Ctx) error {
	return c.Status(200).SendString("Product " + c.Params("id"))
}
