package transaction

import (
	"net/http"
	"time"

	"github.com/LulianoM/bank-api/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (th *TransactionHanlder) Create(c *fiber.Ctx) error {
	var newTransaction models.Transaction

	if err := c.BodyParser(&newTransaction); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	newTransaction.ID = uuid.New()
	newTransaction.EventDate = time.Now()

	// validate bodys

	if err := th.transactionController.Create(newTransaction); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(newTransaction)
}
