package interfaces

import "github.com/LulianoM/bank-api/src/models"

type TransactionController interface {
	Create(models.Transaction) error
}

type TransactionRepositories interface {
	Create(models.Transaction) error
}
