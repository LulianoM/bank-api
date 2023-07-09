package interfaces

import (
	"github.com/LulianoM/bank-api/src/models"
	"github.com/google/uuid"
)

type AccountController interface {
	Create(models.Account) error
	GetByID(uuid.UUID) (models.Account, error)
}

type AccountRepositories interface {
	Create(models.Account) error
	GetByID(uuid.UUID) (models.Account, error)
}
