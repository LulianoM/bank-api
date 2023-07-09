package controllers

import (
	"github.com/LulianoM/bank-api/src/controllers/account"
	"github.com/LulianoM/bank-api/src/interfaces"
	"github.com/LulianoM/bank-api/src/repositories"
)

type ControllersContainer struct {
	AccountController interfaces.AccountController
}

func NewControllerContainer(repositories *repositories.RepositoriesContainer) *ControllersContainer {
	controllersContainer := &ControllersContainer{}

	controllersContainer.AccountController = account.NewAccountController(repositories)

	return controllersContainer
}
