package account

import (
	"fmt"
	"net/http"
	"time"

	"github.com/LulianoM/bank-api/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (ah *AccountHanlder) Create(c *fiber.Ctx) error {
	var newAccount models.Account

	if err := c.BodyParser(&newAccount); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	// TECH DEBT: Validate body

	newAccount.ID = uuid.New()
	newAccount.EventDate = time.Now()

	if err := ah.accountController.Create(newAccount); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}
	return c.Status(http.StatusOK).JSON(newAccount)
}

func (ah *AccountHanlder) Get(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("uuid error: %s", err.Error()),
		})
	}

	account, err := ah.accountController.GetByID(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(account)
}
