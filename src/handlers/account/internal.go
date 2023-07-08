package account

import (
	"net/http"

	"github.com/LulianoM/bank-api/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (ah *AccountHanlder) Create(c *fiber.Ctx) error {
	var newAccount models.Account

	if err := c.BodyParser(&newAccount); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	newAccount.ID = uuid.New()

	if err := ah.accountController.Create(newAccount); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}
	return c.Status(http.StatusOK).JSON(newAccount)
}
