package routes

import (
	"GOHR/controllers"

	"github.com/gofiber/fiber/v2"
)

// SetupChannelRoutes sets up the channel routes
func SetupChannelRoutes(app *fiber.App) {
	router := app.Group("/channels")

	// Re-route into other routers
	router.Use("/:channelId/posts", controllers.GetPosts)

	// Define the route handlers
	router.Get("/", controllers.GetChannels)
	router.Post("/", controllers.CreateChannel)
	router.Get("/:id", controllers.GetChannel)
	router.Put("/:id", controllers.UpdateChannel)
	router.Delete("/:id", controllers.DeleteChannel)
}
