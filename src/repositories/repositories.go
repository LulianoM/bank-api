package repositories

import (
	"github.com/LulianoM/bank-api/src/interfaces"
	"github.com/LulianoM/bank-api/src/repositories/account"
	"github.com/LulianoM/bank-api/src/repositories/credit"
	"github.com/LulianoM/bank-api/src/repositories/transaction"
)

type RepositoriesContainer struct {
	AccountRepositories     interfaces.AccountRepositories
	TransactionRepositories interfaces.TransactionController
	CreditRepositories      interfaces.CreditRepositories
}

func NewRepositoriesContainer() *RepositoriesContainer {
	repositoriesContainer := &RepositoriesContainer{}

	repositoriesContainer.AccountRepositories = account.NewAccountRepositories()
	repositoriesContainer.TransactionRepositories = transaction.NewTransactionRepositories()
	repositoriesContainer.CreditRepositories = credit.NewCreditRepositories()

	return repositoriesContainer
}
