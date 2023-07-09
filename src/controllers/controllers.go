package controllers

import (
	"github.com/LulianoM/bank-api/src/controllers/account"
	"github.com/LulianoM/bank-api/src/controllers/transaction"
	"github.com/LulianoM/bank-api/src/interfaces"
	"github.com/LulianoM/bank-api/src/repositories"
)

type ControllersContainer struct {
	AccountController     interfaces.AccountController
	TransactionController interfaces.TransactionController
}

func NewControllerContainer(repositories *repositories.RepositoriesContainer) *ControllersContainer {
	controllersContainer := &ControllersContainer{}

	controllersContainer.AccountController = account.NewAccountController(repositories)
	controllersContainer.TransactionController = transaction.NewTransactionController(repositories)

	return controllersContainer
}
