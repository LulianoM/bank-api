package account

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (ah *AccountHanlder) Create(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("Alive")
}
