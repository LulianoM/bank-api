package models

import (
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
	re := regexp.MustCompile(`^\d{3}\.\d{3}\.\d{3}-\d{2}$`)
	return re.MatchString(a.DocumentNumber)
}
