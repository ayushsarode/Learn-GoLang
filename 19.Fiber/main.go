package main

import "github.com/gofiber/fiber/v2"

func main (){
	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World from ayush!")
		
	// })

	// app.Get("/:value", func (c *fiber.Ctx) error  {
	// 	return c.SendString("Value: " + c.Params("value"))
	// })

	app.Get("/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("Hello " + c.Params("name"))
			// => Hello john
		}
		return c.SendString("Where is you?")
	})

	app.Listen(":3000")
}