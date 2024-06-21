package handlers

import (
	"test/internal/middleware"
	"test/internal/repository"
	"test/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

func ListNews(c *fiber.Ctx) error {
	page, limit, err := middleware.PaginationParams(c)
	if err != nil {
		return err
	}

	news, err := repository.GetNewsList(page, limit)
	if err != nil {
		logger.Error(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get news list"})
	}

	return c.JSON(fiber.Map{
		"Success": true,
		"News":    news,
	})
}
