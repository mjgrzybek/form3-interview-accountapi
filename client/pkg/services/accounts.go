package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"path"
	"strconv"

	internal "github.com/mjgrzybek/form3-interview-accountapi/client/internal/client"
	"github.com/mjgrzybek/form3-interview-accountapi/client/pkg/models"
)

type AccountsApiService internal.Client

func (svc AccountsApiService) path() string {
	return svc.ApiUrl + "/" + path.Join("organisation", "accounts") // TODO: make it smarter
}

func NewAccountsApiService() *AccountsApiService {
	return (*AccountsApiService)(internal.NewClient())
}

func (svc AccountsApiService) Create(accountData *models.AccountData) (*http.Response, error) {
	buffer, err := encode(models.AccountDataRequest{Data: accountData})
	if err != nil {
		return nil, err
	}

	return svc.HttpClient.Post(svc.path(), "application/vnd.api+json", buffer)
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

func (svc AccountsApiService) Delete(data *models.AccountData) (*http.Response, error) {
	url, err := url.Parse(svc.path() + "/" + data.ID)
	if err != nil {
		return nil, err
	}
	values := url.Query()
	if data.Version != nil {
		values.Add("version", strconv.FormatInt(*data.Version, 10))
	}

	url.RawQuery = values.Encode()

	req, err := http.NewRequest("DELETE", url.String(), bytes.NewReader(nil))
	if err != nil {
		return nil, err
	}
	httpResponse, err := svc.HttpClient.Do(req)
	if err != nil {
		return httpResponse, err
	}
	if httpResponse.StatusCode == http.StatusNoContent {
		return httpResponse, nil
	}

	var errorResponse models.ErrorResponse
	err = json.NewDecoder(httpResponse.Body).Decode(&errorResponse)
	if err != nil {
		return httpResponse, err
	}

	return httpResponse, errors.New(errorResponse.ErrorMessage)
}
