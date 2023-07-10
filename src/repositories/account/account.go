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

func (ar *AccountRepositories) AccountExists(id uuid.UUID) bool {
	var account models.Account
	var count int64

	database.DB.First(&account, "id = ?", id).Count(&count)

	return count > 0
}
