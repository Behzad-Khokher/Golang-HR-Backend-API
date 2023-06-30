package controllers

import (
	"GOHR/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetPosts returns posts
// @desc      Get posts
// @route     GET /api/v1/posts
// @route     GET /api/v1/channels/:channelId/posts
// @access    Public
func GetPosts(c *fiber.Ctx) error {
	c_ID := c.Params("channelId")
	channelID, err := strconv.Atoi(c_ID)

	if err != nil {
		return fiber.DefaultErrorHandler(c, err)
	}

	if c_ID != "" {
		posts := []models.Post{
			{ID: 1, Title: "Post 1", ChannelID: channelID},
			{ID: 2, Title: "Post 2", ChannelID: channelID},
		}

		return c.JSON(posts)
	}

	posts := []models.Post{
		{ID: 1, Title: "Post 1", ChannelID: 1},
		{ID: 2, Title: "Post 2", ChannelID: 2},
	}

	return c.JSON(posts)
}
