package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ruiborda/go-mongodb-crud-api-rest/api/products"
	"github.com/ruiborda/go-mongodb-crud-api-rest/middleware"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "golang-fiber",
		AppName:       "Bar",
	})

	app.Use(middleware.SecurityHeaders)

	app.Static("/", "./public")

	app.Mount("/products", products.Routes())

	app.Get("/:name", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html; charset=utf-8")
		return c.SendString("Hello " + c.Params("name"))
	})
	err := app.Listen(":3000")

	if err != nil {
		fmt.Println(err)
		return
	}
}
