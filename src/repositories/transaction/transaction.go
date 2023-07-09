package transaction

import (
	"github.com/LulianoM/bank-api/internal/database"
	"github.com/LulianoM/bank-api/src/models"
)

type TransactionRepositories struct{}

func NewTransactionRepositories() *TransactionRepositories {
	return &TransactionRepositories{}
}

func (tr *TransactionRepositories) Create(transaction models.Transaction) error {
	if result := database.DB.Create(&transaction); result.Error != nil {
		return result.Error
	}
	return nil
}
