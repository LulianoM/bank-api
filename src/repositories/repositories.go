package repositories

import (
	"github.com/LulianoM/bank-api/src/interfaces"
	"github.com/LulianoM/bank-api/src/repositories/account"
	"github.com/LulianoM/bank-api/src/repositories/transaction"
)

type RepositoriesContainer struct {
	AccountRepositories     interfaces.AccountRepositories
	TransactionRepositories interfaces.TransactionController
}

func NewRepositoriesContainer() *RepositoriesContainer {
	repositoriesContainer := &RepositoriesContainer{}

	repositoriesContainer.AccountRepositories = account.NewAccountRepositories()
	repositoriesContainer.TransactionRepositories = transaction.NewTransactionRepositories()

	return repositoriesContainer
}
