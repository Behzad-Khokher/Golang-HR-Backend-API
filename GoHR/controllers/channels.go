package controllers

import (
	"GOHR/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetChannels returns all channels
// @desc      Get all channels
// @route     GET /api/v1/channels
// @access    Public
func GetChannels(c *fiber.Ctx) error {
	channels := []models.Channel{
		{ID: 1, Name: "Channel 1", Description: "Channel 1 description", PostCount: 5},
		{ID: 2, Name: "Channel 2", Description: "Channel 2 description", PostCount: 3},
	}

	return c.JSON(channels)
}

// GetChannel returns a single channel by ID
// @desc      Get single channel
// @route     GET /api/v1/channel/:id
// @access    Public
func GetChannel(c *fiber.Ctx) error {
	id := c.Params("id")
	channelID, err := strconv.Atoi(id)

	if err != nil {
		return fiber.DefaultErrorHandler(c, err)
	}

	channel := models.Channel{ID: channelID, Name: "Channel", Description: "Channel description", PostCount: 5}

	return c.JSON(channel)
}

// CreateChannel creates a new channel
// @desc      Create new channel
// @route     POST /api/v1/channels
// @access    Private
func CreateChannel(c *fiber.Ctx) error {
	newChannel := models.Channel{ID: 1, Name: "New Channel", Description: "New Channel description", PostCount: 0}

	return c.JSON(newChannel)
}

// UpdateChannel updates an existing channel by ID
// @desc      Update channel
// @route     PUT /api/v1/channels/:id
// @access    Private
func UpdateChannel(c *fiber.Ctx) error {
	id := c.Params("id")
	channelID, err := strconv.Atoi(id)

	if err != nil {
		return fiber.DefaultErrorHandler(c, err)
	}

	updatedChannel := models.Channel{ID: channelID, Name: "Updated Channel", Description: "Updated Channel description", PostCount: 5}

	return c.JSON(updatedChannel)
}

// DeleteChannel deletes a channel by ID
// @desc      Delete channel
// @route     DELETE /api/v1/channels/:id
// @access    Private
func DeleteChannel(c *fiber.Ctx) error {
	return c.SendString("Channel Removed")
}

// ChannelPhotoUpload handles photo upload for a channel
// @desc      Upload photo for channel
// @route     PUT /api/v1/channels/:id/photo
// @access    Private
func ChannelPhotoUpload(c *fiber.Ctx) error {
	return c.SendString("Upload Photo for Channel")
}
