package types

import (
	"time"

	"github.com/google/uuid"
)

type (
	Account struct {
		AccountNumber string    `json:"accountNumber"`
		FirstName     string    `json:"firstName"`
		LastName      string    `json:"lastName"`
		Email         string    `json:"email"`
		Balance       int64     `json:"balance"`
		CreatedAt     time.Time `json:"createdAt"`
	}

	AccountRequest struct {
		FirstName string `json:"firstName" validate:"required"`
		LastName  string `json:"lastName" validate:"required"`
		Email     string `json:"email" validate:"required,email"`
	}
)

func NewAccount(firstName, lastName, email string) *Account {
	return &Account{
		AccountNumber: uuid.NewString(),
		FirstName:     firstName,
		LastName:      lastName,
		Email:         email,
		CreatedAt:     time.Now().UTC(),
	}
}
