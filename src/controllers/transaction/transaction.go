package transaction

import (
	"errors"

	"github.com/LulianoM/bank-api/src/interfaces"
	"github.com/LulianoM/bank-api/src/models"
	"github.com/LulianoM/bank-api/src/repositories"
)

type TransactionController struct {
	transactionRepositories interfaces.TransactionRepositories
	accountRepositories     interfaces.AccountRepositories
}

func NewTransactionController(repositories *repositories.RepositoriesContainer) *TransactionController {
	return &TransactionController{
		transactionRepositories: repositories.TransactionRepositories,
		accountRepositories:     repositories.AccountRepositories,
	}
}

func (tc *TransactionController) Create(transaction models.Transaction) error {

	if !tc.accountRepositories.AccountExists(transaction.AccountID) {
		return errors.New("account dont exists")
	}

	return tc.transactionRepositories.Create(transaction)
}
