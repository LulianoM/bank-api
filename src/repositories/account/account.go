package account

import (
	"github.com/LulianoM/bank-api/internal/database"
	"github.com/LulianoM/bank-api/src/models"
	"github.com/google/uuid"
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

func (ar *AccountRepositories) GetByID(id uuid.UUID) (models.Account, error) {
	var account models.Account
	if result := database.DB.First(&account, "id = ?", id); result.Error != nil {
		return account, result.Error
	}
	return account, nil
}
