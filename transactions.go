package main

import (
	"github.com/labstack/echo/v4"
)

type TransactionService struct {
  store Store
}

func NewTransactionService(s Store) *TransactionService {
  return &TransactionService{store:  s}
}

func (s *TransactionService) RegisterRoutes(e *echo.Echo) {
  e.GET("/transaction/{id}", handleGetTransaction)
  e.POST("/transaction", handleCreateTransaction)
}

func handleGetTransaction(c echo.Context) error {
  return nil
}
func handleCreateTransaction(c echo.Context) error {
  return nil
}
