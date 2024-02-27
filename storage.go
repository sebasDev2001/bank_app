package main

import "database/sql"

type Store interface {
  CreateAccount() error 
  CreateTransaction() error
}

type Storage struct {
  db *sql.DB
}

func NewStore(db *sql.DB) *Storage {
  return &Storage{db: db}
}

func (s *Storage) CreateAccount() error {
  return nil
}

func (s *Storage) CreateTransaction() error {
  return nil
}
