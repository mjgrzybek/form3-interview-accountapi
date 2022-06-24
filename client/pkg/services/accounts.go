package client

import (
	internal "github.com/mjgrzybek/form3-interview-accountapi/client/internal/client"
	"github.com/mjgrzybek/form3-interview-accountapi/client/pkg/models"
)

type AccountsApiService internal.Client

func NewAccountsApiService() *AccountsApiService {
	return (*AccountsApiService)(internal.NewClient())
}

func (svc AccountsApiService) Create(*models.AccountData) error {
	return nil
}

func (svc AccountsApiService) Fetch() error {
	return nil
}

func (svc AccountsApiService) Delete() error {
	return nil
}
