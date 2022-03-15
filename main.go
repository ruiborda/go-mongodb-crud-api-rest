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

	err := app.Listen(":3000")

	if err != nil {
		fmt.Println(err)
		return
	}
}
