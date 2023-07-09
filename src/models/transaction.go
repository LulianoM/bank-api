package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Metadata

	Account   *Account  `json:"-"`
	AccountID uuid.UUID `json:"acount_id"`

	ID          uuid.UUID `json:"id" gorm:"unique"`
	OperationID int       `json:"operation_id"`
	Amount      float64   `json:"amount"`
	EventDate   time.Time `json:"event_date"`
}
