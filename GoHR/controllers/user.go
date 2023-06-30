/*
Note: User Controller Under Development
Returns stub for testing
*/
package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// GetUsers returns a list of users
func GetUsers(c *fiber.Ctx) error {
	return c.SendString("Get Users")
}

// GetUser returns a single user by ID
func GetUser(c *fiber.Ctx) error {
	return c.SendString("Get User")
}

// CreateUser creates a new user
func CreateUser(c *fiber.Ctx) error {
	return c.SendString("Create User")
}

// UpdateUser updates an existing user by ID
func UpdateUser(c *fiber.Ctx) error {
	return c.SendString("Update User")
}

// DeleteUser deletes a user by ID
func DeleteUser(c *fiber.Ctx) error {
	return c.SendString("Delete User")
}
