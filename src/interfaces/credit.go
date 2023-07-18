package interfaces

import (
	"github.com/LulianoM/bank-api/src/models"
	"github.com/google/uuid"
)

type CreditRepositories interface {
	Create(credit models.Credit) error
	GetAvalableCreditLimit(uuid.UUID) (float64, error)
	UpdatedAvalableCreditLimit(uuid.UUID, float64) error
}
