package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/sebasdev2001/bank_app/types"
)

type AccountSerivce struct {
	store Store
}

func NewAccountService(s Store) *AccountSerivce {
	return &AccountSerivce{store: s}
}

func (s *AccountSerivce) RegisterRoutes(e *echo.Echo) {
	e.GET("/account/:id", s.handleGetAccount)
	e.POST("/account", s.handleCreateAccount)
  e.PATCH("/account/:id", s.handleUpdateAccount)
	e.DELETE("/account/:id", s.handleDeleteAccount)
}

func (s *AccountSerivce) handleGetAccount(c echo.Context) error {
	account_id := c.Param("id")
	if err := uuid.Validate(account_id); err != nil {
		return err // change this to be a json response with error
	}
	fmt.Printf("account_id: %v\n", account_id)
  
  acc, err := s.store.GetAccount(account_id)

  if err != nil {
    return err
  }

	return c.JSON(http.StatusOK, &acc)
}

func (s *AccountSerivce) handleCreateAccount(c echo.Context) error {
	accRequest := new(types.AccountRequest)
	if err := c.Bind(accRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(accRequest); err != nil {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{Error: fmt.Sprintf("Error in account validation: %v", err)})
	}
	acc := types.NewAccount(accRequest.FirstName, accRequest.LastName, accRequest.Email)

  if err:= s.store.CreateAccount(acc); err != nil {
    return c.JSON(http.StatusBadRequest, &types.ErrorResponse{Error: fmt.Sprintf("%v", err.Error())})
  }

	return c.JSON(http.StatusOK, acc)
}

func (s *AccountSerivce) handleDeleteAccount(c echo.Context) error {
	return nil
}

func (s *AccountSerivce) handleUpdateAccount(c echo.Context) error {
  account_id := c.Param("id")
  if err := uuid.Validate(account_id); err != nil {
    return err
  }
  fmt.Printf("account_id: %v\n", account_id)

  return c.JSON(http.StatusOK, &types.Account{})
}
