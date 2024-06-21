package middleware

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ValidateRequest(c *fiber.Ctx) error {
	body := c.Body()
	if len(body) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request body is empty"})
	}
	return nil
}

func PaginationParams(c *fiber.Ctx) (int, int, error) {
	pageStr := c.Query("page", "1")
	limitStr := c.Query("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return 0, 0, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid page parameter"})
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return 0, 0, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid limit parameter"})
	}
	return page, limit, nil
}
