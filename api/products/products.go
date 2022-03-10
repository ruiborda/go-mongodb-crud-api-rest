package products

import "github.com/gofiber/fiber/v2"

func Routes() *fiber.App {
	products := fiber.New()

	products.Get("/", IndexProducts)
	products.Post("/", StoreProducts)
	products.Get("/:id", ShowProducts)
	products.Put("/:id", UpdateProducts)
	products.Delete("/:id", DeleteProducts)

	return products
}
