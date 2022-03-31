package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Healthcheck(c *fiber.Ctx) error {

	return c.Status(http.StatusOK).JSON(&fiber.Map{"message": "ok"})
}
