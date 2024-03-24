package main

import (
	"database/sql"
	"fmt"

	"github.com/sebasdev2001/bank_app/types"
)

type Store interface {
	CreateAccount(*types.Account) error
  GetAccount(string) (*types.Account, error)
	DeleteAccount(int) error
	UpdateAccount(*types.Account) error
	CreateTransaction(*types.Transaction) error
}

type Storage struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) CreateAccount(acc *types.Account) error {

  _, err :=  s.GetAccount(acc.Email)
  if err != nil {
    return err
  }

	query := `insert into account
  (first_name, last_name, email, balance, created_at)
  values ($1, $2, $3, $4, $5);`

	resp, err := s.db.Query(
		query,
		acc.FirstName,
		acc.LastName,
		acc.Email,
		acc.Balance,
		acc.CreatedAt,
	)
	if err != nil {
		return err
	}

	fmt.Printf("Account creation resp: %v\n", resp)

	return nil
}

func (s *Storage) GetAccount(email string) (*types.Account, error) {
  return nil, nil
}

func (s *Storage) DeleteAccount(id int) error {
	return nil
}

func (s *Storage) UpdateAccount(acc *types.Account) error {
	return nil
}

func (s *Storage) CreateTransaction(transaction *types.Transaction) error {
	return nil
}
