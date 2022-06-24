package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"path"

	internal "github.com/mjgrzybek/form3-interview-accountapi/client/internal/client"
)

type AccountsApiService internal.Client

func NewAccountsApiService() *AccountsApiService {
	return (*AccountsApiService)(internal.NewClient())
}

func (svc AccountsApiService) Create(accountData any) (*http.Response, error) {
	url := svc.ApiUrl + "/" + path.Join("organisation", "accounts") // TODO: make it smarter

	buffer, err := encode(accountData)
	if err != nil {
		return nil, err
	}

	return svc.HttpClient.Post(url, "application/vnd.api+json", buffer)
}

func encode(data any) (*bytes.Buffer, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(body), nil
}

func (svc AccountsApiService) Fetch() error {
	return nil
}

func (svc AccountsApiService) Delete() error {
	return nil
}
