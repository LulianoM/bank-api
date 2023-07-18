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
	creditRepositories      interfaces.CreditRepositories
}

func NewTransactionController(repositories *repositories.RepositoriesContainer) *TransactionController {
	return &TransactionController{
		transactionRepositories: repositories.TransactionRepositories,
		accountRepositories:     repositories.AccountRepositories,
		creditRepositories:      repositories.CreditRepositories,
	}
}

func (tc *TransactionController) Create(transaction models.Transaction) error {

	if !tc.accountRepositories.AccountExists(transaction.AccountID) {
		return errors.New("account dont exists")
	}

	avalableCredit, err := tc.creditRepositories.GetAvalableCreditLimit(transaction.AccountID)
	if err != nil {
		return err
	}

	if !isAvalibleCreditLimitToTransaction(transaction.Amount, avalableCredit) {
		return errors.New("no credit avalable")
	}

	if err := tc.transactionRepositories.Create(transaction); err != nil {
		return err
	}

	newAmountValue := (avalableCredit + transaction.Amount)

	if err := tc.creditRepositories.UpdatedAvalableCreditLimit(transaction.AccountID, newAmountValue); err != nil {
		return err
	}

	return nil
}

func isAvalibleCreditLimitToTransaction(amount float64, avalableCredit float64) bool {
	if (avalableCredit - amount) < 0 {
		return false
	}
	return true
}
