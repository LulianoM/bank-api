package account

import (
	"github.com/LulianoM/bank-api/src/interfaces"
	"github.com/LulianoM/bank-api/src/models"
	"github.com/LulianoM/bank-api/src/repositories"
	"github.com/google/uuid"
)

type AccountController struct {
	accountRepositories interfaces.AccountRepositories
	creditRepositories  interfaces.CreditRepositories
}

func NewAccountController(repositories *repositories.RepositoriesContainer) *AccountController {
	return &AccountController{accountRepositories: repositories.AccountRepositories,
		creditRepositories: repositories.CreditRepositories,
	}
}

func (ac *AccountController) Create(account models.Account) error {
	var newCredit models.Credit
	newCredit.NewCredit(account.ID)

	if err := ac.creditRepositories.Create(newCredit); err != nil {
		return err
	}

	return ac.accountRepositories.Create(account)
}

func (ac *AccountController) GetByID(id uuid.UUID) (models.Account, error) {
	return ac.accountRepositories.GetByID(id)
}
