package controllers

import (
	"github.com/LulianoM/bank-api/src/controllers/account"
	"github.com/LulianoM/bank-api/src/interfaces"
)

type ControllersContainer struct {
	AccountController interfaces.AccountController
}

func NewControllerContainer() *ControllersContainer {
	controllersContainer := &ControllersContainer{}

	controllersContainer.AccountController = account.NewAccountController()

	return controllersContainer
}
