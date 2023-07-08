package account

import (
	"github.com/LulianoM/bank-api/src/models"
)

type AccountController struct{}

func NewAccountController() *AccountController {
	return &AccountController{}
}

func (ac *AccountController) Create(account models.Account) error {
	return nil
}
