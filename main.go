package main

import (
	"log"

	"github.com/dhruvdabhi101/gurl-shortner/database"
	"github.com/dhruvdabhi101/gurl-shortner/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CreateServer() *fiber.App {
  app := fiber.New()

  return app
}


func main() {
  database.ConnectToDB()
  app := CreateServer()

  app.Use(cors.New())

  router.SetupRoutes(app)


  app.Use(func (c *fiber.Ctx) error {
    return c.SendStatus(404)
  })

  log.Fatal(app.Listen(":3000"))
}
