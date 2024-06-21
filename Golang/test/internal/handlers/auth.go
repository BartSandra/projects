package handlers

import (
	"github.com/gofiber/fiber/v2"
	"test/pkg/utils"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	var creds Credentials
	if err := c.BodyParser(&creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Ошибка при разборе данных"})
	}

	if creds.Username != "expected_username" || creds.Password != "expected_password" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Неверные учетные данные"})
	}

	token, err := utils.GenerateJWT(1)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Не удалось сгенерировать токен"})
	}

	return c.JSON(fiber.Map{"token": token})
}