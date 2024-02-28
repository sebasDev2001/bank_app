package types

import (
	"time"

	"github.com/google/uuid"
)

type (
	Transaction struct {
		ID        string    `json:"id"`
		FromUser  string    `json:"fromUser"` // uuid
		ToUser    string    `json:"toUser"`   // uuid
		CreatedAt time.Time `json:"created_at"`
		Amount    int64     `json:"amount"`
	}
)

func NewTransaction(from, to *Account, amount int64) *Transaction {
	return &Transaction{
		ID:        uuid.NewString(),
		FromUser:  from.AccountNumber,
		ToUser:    to.AccountNumber,
		CreatedAt: time.Now(),
		Amount:    amount,
	}
}
