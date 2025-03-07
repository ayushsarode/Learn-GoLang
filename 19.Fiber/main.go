package main

import (
	"encoding/json"
	"fmt"
    "github.com/gofiber/fiber/v2/middleware/basicauth"

	"github.com/gofiber/fiber/v2"
)
var apiHandler = func(c *fiber.Ctx) error {return nil}


func main (){
	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World from ayush!")
		
	// })

	app.Static("/", "./public")

	// app.Get("/:value", func (c *fiber.Ctx) error  {
	// 	return c.SendString("Value: " + c.Params("value"))
	// })

	// app.Get("/:name?", func(c *fiber.Ctx) error {
	// 	if c.Params("name") != "" {
	// 		return c.SendString("Hello " + c.Params("name"))
	// 		// => Hello john
	// 	}
	// 	return c.SendString("Where is you?")
	// })


	// app.Static("/", "./public/index.html", fiber.Static{
	// 	Compress:      true,
	// 	ByteRange:     true,
	// 	Browse:        true,
	// 	Index:         "john.html",
	// 	CacheDuration: 10 * time.Second,
	// 	MaxAge:        3600,
	// })
	// app.Static("/about", "./public/about.html")

	app.Get("/api/list", func (c *fiber.Ctx)  error{
		return c.SendString("This is Get Req")
	})

	app.Post("/api/register", func(c *fiber.Ctx) error {
		return c.SendString("I'm a POST request!")
	  })

	  

	//   app.Route("/test", func(api fiber.Router) {
	// 	handler := nil
	// 	api.Get("/foo", handler) // /test/foo (name: test.foo)
	// 	api.Get("/bar", handler).Name("bar") // /test/bar (name: test.bar)
	// }, "test.")



	app.Get("/ayush", apiHandler)
    data, _ := json.MarshalIndent(app.Stack(), "", "  ")
    fmt.Println(string(data))


	app.Get("/", apiHandler).Name("index")
    
    data, _ = json.MarshalIndent(app.GetRoute("index"), "", "  ")
	fmt.Print(string(data))

	app.Get("/ip", func (c *fiber.Ctx)  error {
		return c.SendString(c.IP());
	})

	app.Get("/queries", func (c *fiber.Ctx) error  {
		m := c.Queries()
		first := m["first"]
		second := m["second"]

		return c.JSON(fiber.Map {
			"first": first,
			"second": second,
		})
	})

	app.Get("/fiber", func(c *fiber.Ctx) error {
		c.Status(fiber.StatusOK) // Set HTTP 200 status
		return c.SendString("Fiber is running!") // Send response 
	})
	


	app.Listen(":3000")
}