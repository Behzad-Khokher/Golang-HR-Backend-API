package routes

import (
	"GOHR/controllers"

	"github.com/gofiber/fiber/v2"
)

// SetupUserRoutes sets up the user routes
func SetupUserRoutes(app *fiber.App) {
	userRouter := app.Group("/users")

	userRouter.Get("/", controllers.GetUsers)
	userRouter.Post("/", controllers.CreateUser)
	userRouter.Get("/:id", controllers.GetUser)
	userRouter.Put("/:id", controllers.UpdateUser)
	userRouter.Delete("/:id", controllers.DeleteUser)
}
