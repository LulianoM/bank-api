package models

import (
	"github.com/google/uuid"
)

type Credit struct {
	Metadata

	Account   *Account  `json:"-"`
	AccountID uuid.UUID `json:"account_id"`

	ID                  uuid.UUID `json:"id" gorm:"unique"`
	AvalableCreditLimit float64   `json:"avalable_credit_limit"`
}

func (c *Credit) NewCredit(id uuid.UUID) {
	c.AccountID = id
	c.AvalableCreditLimit = 0
	c.ID = uuid.New()
}
