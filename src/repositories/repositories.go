package repositories

import (
	"github.com/LulianoM/bank-api/src/interfaces"
	"github.com/LulianoM/bank-api/src/repositories/account"
)

type RepositoriesContainer struct {
	AccountRepositories interfaces.AccountRepositories
}

func NewRepositoriesContainer() *RepositoriesContainer {
	repositoriesContainer := &RepositoriesContainer{}

	repositoriesContainer.AccountRepositories = account.NewAccountRepositories()

	return repositoriesContainer
}
