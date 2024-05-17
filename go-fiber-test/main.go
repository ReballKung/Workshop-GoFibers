package main

import (
	router "go-fiber-test/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	router.InetRoutes(app)
	app.Listen(":3000")
}
