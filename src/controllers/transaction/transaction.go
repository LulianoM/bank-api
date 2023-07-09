package transaction

import (
	"github.com/LulianoM/bank-api/src/interfaces"
	"github.com/LulianoM/bank-api/src/models"
	"github.com/LulianoM/bank-api/src/repositories"
)

type TransactionController struct {
	transactionRepositories interfaces.TransactionRepositories
}

func NewTransactionController(repositories *repositories.RepositoriesContainer) *TransactionController {
	return &TransactionController{transactionRepositories: repositories.TransactionRepositories}
}

func (tc *TransactionController) Create(transaction models.Transaction) error {
	return tc.transactionRepositories.Create(transaction)
}
