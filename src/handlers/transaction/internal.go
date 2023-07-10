package transaction

import (
	"net/http"

	"github.com/LulianoM/bank-api/internal/httputils"
	"github.com/LulianoM/bank-api/src/models"
	"github.com/gofiber/fiber/v2"
)

func (th *TransactionHanlder) Create(c *fiber.Ctx) error {
	var newTransaction models.Transaction

	if err := c.BodyParser(&newTransaction); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httputils.BuildErrorResponse(err))
	}

	if err := newTransaction.ValidateBody(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httputils.BuildErrorResponse(err))
	}

	newTransaction.SetTransactionData()

	newTransaction.SetTransactionAmount()

	if err := th.transactionController.Create(newTransaction); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httputils.BuildErrorResponse(err))
	}

	return c.Status(http.StatusOK).JSON(httputils.BuildItemResponse(newTransaction))
}
