package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

)

type APIServer struct {
	address string
	store Store
}

func NewAPIServer(address string, store Store) *APIServer {
	return &APIServer{
		address: address,
		store:   store,
	}
}

func (api *APIServer) Run() {
  AccountService := NewAccountSertice(api.store)
  TransactionService := NewTransactionService(api.store)
	e := echo.New()
	e.Use(middleware.Logger())

  AccountService.RegisterRoutes(e)
  go TransactionService.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(api.address))
}
