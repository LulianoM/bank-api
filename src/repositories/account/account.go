package account

import (
	"github.com/LulianoM/bank-api/internal/database"
	"github.com/LulianoM/bank-api/src/models"
)

type AccountRepositories struct{}

func NewAccountRepositories() *AccountRepositories {
	return &AccountRepositories{}
}

func (ar *AccountRepositories) Create(account models.Account) error {
	if result := database.DB.Create(&account); result.Error != nil {
		return result.Error
	}
	return nil
}
