package account

import (
	"github.com/LulianoM/bank-api/src/interfaces"
	"github.com/LulianoM/bank-api/src/models"
	"github.com/LulianoM/bank-api/src/repositories"
)

type AccountController struct {
	accountRepositoes interfaces.AccountRepositories
}

func NewAccountController(repositories *repositories.RepositoriesContainer) *AccountController {
	return &AccountController{accountRepositoes: repositories.AccountRepositories}
}

func (ac *AccountController) Create(account models.Account) error {
	return ac.accountRepositoes.Create(account)
}
