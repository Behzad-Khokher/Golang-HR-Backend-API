package main

import (
	"GOHR/database"
	"GOHR/routes"

	"log"

	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	fmt.Println("Go Lang Server Running")
	app := fiber.New()

	// Connect Database, Load Models to postgress: *Note: Modify database/database.go so connection arguments can be passed from main for single point of access
	database.ConnectDb()

	// Pre-built Middleware: *Note: Need to set up custom config files
	// app.Use(logger.New())
	app.Use(recover.New())
	app.Use(limiter.New())
	app.Use(cors.New())
	app.Use(helmet.New())

	// Custome Loggers: * Note: Seperate middleware loggers in a seprate file

	// Loggers middleware to log every request to command line:
	app.Use(func(c *fiber.Ctx) error {
		log.Printf("Request: %s %s\n", c.Method(), c.Path())
		return c.Next()
	})

	// Middleware to handle wrong routes:
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Not Found",
		})
	})

	// Mount Routers: * Note: Create a single point of access for all routes
	// * Note: Create app group here for: /api/v1 -> which will point to to other routes
	// Router := app.Group("/api/v1/", <point to to other routes here>)

	// Temporary direct route call for testing, *Make route group here:
	routes.SetupUserRoutes(app)
	routes.SetupChannelRoutes(app)
	routes.SetupPostRoutes(app)

	app.Listen(":3000")

}
