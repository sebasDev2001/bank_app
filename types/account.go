package types

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type (
	Account struct {
		ID            int       `json:"id"`
		AccountNumber string    `json:"accountNumber"`
		FirstName     string    `json:"firstName"`
		LastName      string    `json:"lastName"`
		Email         string    `json:"email"`
		Balance       int64     `json:"balance"`
		CreatedAt     time.Time `json:"createdAt"`
	}

  CreateAccountRequest struct {
    FirstName string `json:"firstName"`
    LastName string `json:"lastName"`
    Email string `json:"email"`
  }

)

func NewAccount(firstName, lastName, email string) *Account {
	return &Account{
		ID:            rand.Intn(100000),
		AccountNumber: uuid.NewString(),
		FirstName:     firstName,
		LastName:      lastName,
		Email:         email,
		CreatedAt:     time.Now().UTC(),
	}
}

