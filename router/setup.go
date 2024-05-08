package router

import "github.com/gofiber/fiber/v2"

var USER fiber.Router

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	USER = app.Group("/user")

	api.Get("/", hello)

  SetupUserRoutes()


}
