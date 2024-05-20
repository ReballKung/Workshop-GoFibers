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

	//CRUD dogs
	dog := v1.Group("/dog")
	dog.Get("", c.GetDogs)
	dog.Get("/all", c.GetALLDogs)
	dog.Get("/filter", c.GetDog)
	dog.Get("/json", c.GetDogsJson)
	dog.Post("/", c.AddDog) // * c => Create
	dog.Put("/:id", c.UpdateDog)
	dog.Get("/history", c.GetDataDelete) // * Ex 7.0.2

	//CRUD company
	company := v1.Group("/company")
	company.Get("", c.GetCompony)            // * View AllData
	company.Post("/", c.AddCompany)          // * ADD company
	company.Put("/:id", c.UpdateCompany)     // * Update company
	company.Delete("/:id", c.DeleteCompany)  // * Delete company
	company.Get("/filter", c.GetComponyByID) // * View Data By ID

}
