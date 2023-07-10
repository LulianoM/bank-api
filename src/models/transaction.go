package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Metadata

	Account   *Account  `json:"-"`
	AccountID uuid.UUID `json:"account_id"`

	ID          uuid.UUID `json:"id" gorm:"unique"`
	OperationID int       `json:"operation_id"`
	Amount      float64   `json:"amount"`
	EventDate   time.Time `json:"event_date"`
}

func (t *Transaction) SetTransactionAmount() {
	if t.OperationID != 4 {
		t.Amount = t.Amount * -1
	}
}

func (t *Transaction) SetTransactionData() {
	t.ID = uuid.New()
	t.EventDate = time.Now()
}

func (t *Transaction) ValidateBody() error {
	if t.AccountID == uuid.Nil {
		return fmt.Errorf("'account_id' cannot be empty")
	}

	if t.OperationID < 0 {
		return fmt.Errorf("'operation_id' must be a positive integer")
	}

	if t.Amount <= 0 {
		return fmt.Errorf("'amount' must be a number greater than zero")
	}

	return nil
}
