package account

import (
	"net/http"
	"time"

	"github.com/LulianoM/bank-api/internal/httputils"
	"github.com/LulianoM/bank-api/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (ah *AccountHanlder) Create(c *fiber.Ctx) error {
	var newAccount models.Account

	if err := c.BodyParser(&newAccount); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	if err := newAccount.IsValidBody(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	newAccount.ID = uuid.New()
	newAccount.EventDate = time.Now()

	if err := ah.accountController.Create(newAccount); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httputils.BuildErrorResponse(err))
	}
	return c.Status(http.StatusOK).JSON(httputils.BuildItemResponse(newAccount))
}

func (ah *AccountHanlder) Get(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(httputils.BuildErrorItemResponse(err, id))
	}

	account, err := ah.accountController.GetByID(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(httputils.BuildErrorResponse(err))
	}

	return c.Status(http.StatusOK).JSON(httputils.BuildItemResponse(account))
}
