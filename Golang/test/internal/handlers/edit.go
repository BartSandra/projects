package handlers

import (
	"test/internal/middleware"
	"test/internal/repository"
	"test/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

func EditNews(c *fiber.Ctx) error {
	if err := middleware.ValidateRequest(c); err != nil {
		return err
	}

	newsID := c.Params("Id")
	newsInput := new(repository.NewsInput)
	if err := c.BodyParser(newsInput); err != nil {
		logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if err := repository.UpdateNews(newsID, newsInput); err != nil {
		logger.Error(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update news"})
	}

	return c.SendStatus(fiber.StatusOK)
}
