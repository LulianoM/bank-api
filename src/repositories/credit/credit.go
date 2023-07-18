package credit

import (
	"github.com/LulianoM/bank-api/internal/database"
	"github.com/LulianoM/bank-api/src/models"
	"github.com/google/uuid"
)

type CreditRepositories struct{}

func NewCreditRepositories() *CreditRepositories {
	return &CreditRepositories{}
}

func (cr *CreditRepositories) Create(credit models.Credit) error {
	if result := database.DB.Create(&credit); result.Error != nil {
		return result.Error
	}
	return nil
}

func (cr *CreditRepositories) GetAvalableCreditLimit(account_id uuid.UUID) (float64, error) {
	var credit models.Credit

	if result := database.DB.First(&credit, "account_id = ?", account_id); result.Error != nil {
		return credit.AvalableCreditLimit, result.Error
	}
	return credit.AvalableCreditLimit, nil
}

func (cr *CreditRepositories) UpdatedAvalableCreditLimit(account_id uuid.UUID, amount float64) error {
	var credit models.Credit

	if result := database.DB.Where(&credit, "account_id = ?", account_id).Update("avalable_credit_limit", amount); result.Error != nil {
		return result.Error
	}

	return nil
}
