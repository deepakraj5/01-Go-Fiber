package main

import "github.com/gofiber/fiber/v2"

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Fiber")
	})

	app.Get("/api/v1/:name", func(c *fiber.Ctx) error {

		name := c.Params("name")

		return c.Status(200).JSON(&fiber.Map{
			"name": name,
		})

	})

	app.Listen(":3000")

}
