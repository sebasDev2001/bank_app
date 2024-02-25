package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sebasdev2001/bank_app/types"
)

type AccountSerivce struct {
  store Store
}

func NewAccountSertice(s Store) *AccountSerivce {
  return &AccountSerivce{store: s}
}

func (s *AccountSerivce) RegisterRoutes(e *echo.Echo) {
  e.GET("/account/:id", handleGetAccount)
  e.POST("/account", handleCreateAccount)
  e.DELETE("/account/:id", handleDeleteAccount)
}

func handleGetAccount(c echo.Context) error {
  account_id := c.Param("id")
  fmt.Printf("account_id: %v\n", account_id)
  return c.JSON(http.StatusOK, &types.Account{})
}

func handleCreateAccount(c echo.Context) error {
  accRequest := new(types.CreateAccountRequest)
  if err := json.NewDecoder(c.Request().Body).Decode(accRequest); err != nil {
    return fmt.Errorf("error decoding request: %v", err)
  }
  acc := types.NewAccount(accRequest.FirstName, accRequest.LastName, accRequest.Email)
  return c.JSON(http.StatusOK, &acc)
}

func handleDeleteAccount(c echo.Context) error {
	return nil
}
