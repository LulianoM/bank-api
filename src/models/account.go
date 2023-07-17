package models

import (
	"fmt"
	"regexp"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Metadata

	ID             uuid.UUID `json:"id" gorm:"primaryKey" `
	DocumentNumber string    `json:"document_number" gorm:"unique"`
	EventDate      time.Time `json:"event_date"`
}

func (a *Account) IsValidDocumentNumber() bool {
	if a.DocumentNumber == "" {
		return false
	}

	regex := regexp.MustCompile(`^\d+$`)
	return regex.MatchString(a.DocumentNumber)
}

func (a *Account) IsValidBody() error {
	if !a.IsValidDocumentNumber() {
		return fmt.Errorf("invalid document number")
	}

	return nil
}
