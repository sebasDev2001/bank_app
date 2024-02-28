package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type DataValidator struct {
  validator *validator.Validate
}

type APIServer struct {
	address string
	store   Store
  dataValidator *DataValidator
}

func (cv *DataValidator) Validate(i interface{}) error {
  return cv.validator.Struct(i)
}

func NewAPIServer(address string, store Store) *APIServer {
	return &APIServer{
		address: address,
		store:   store,
    dataValidator: &DataValidator{validator: validator.New()},
	}
}

func (api *APIServer) Run() {
	AccountService := NewAccountService(api.store)
	TransactionService := NewTransactionService(api.store)
	e := echo.New()
  e.Validator = api.dataValidator
	e.Use(middleware.Logger())

	AccountService.RegisterRoutes(e)
	TransactionService.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(api.address))
}
