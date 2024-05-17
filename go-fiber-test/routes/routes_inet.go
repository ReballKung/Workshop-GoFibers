package routes

import (
	c "go-fiber-test/controllers" // * EX : 5.3

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {
	//* Create routes
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v3 := api.Group("/v3")

	//* EX : 5.0
	//* [Middleware && Basic Authentication]
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"gofiber": "21022566",
		},
	}))

	//* API endpoint
	v1.Get("/", c.HelloTest)                   //  [TEST AUTH]
	v1.Post("/fact/:number", c.InputFactorial) //* EX: 5.1
	v3.Post("/:name", c.CornvertAscii)         //* EX: 5.2
	v1.Post("/register", c.RegisterEmployee)   //* EX: 6

}
