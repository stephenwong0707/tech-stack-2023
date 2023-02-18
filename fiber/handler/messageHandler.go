package handler

import (
	"github.com/gofiber/fiber/v2"
)

func PostMessage(c *fiber.Ctx) error {
	message, err := messageService.SaveMessage(c.Query("text"))
	if err != nil {
		return c.SendStatus(400)
	}
	return c.JSON(fiber.Map{
		"createTime": message.CreateTime,
		"content":    message.Content,
	})
}
