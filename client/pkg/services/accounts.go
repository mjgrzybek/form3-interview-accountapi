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

func NewAccountsApiService() (*AccountsApiService, error) {
	svc := (*AccountsApiService)(internal.NewClient())

	parsedUrl, err := url.Parse(svc.ApiUrl.String() + path.Join("/", "organisation", "accounts")) // TODO: make it smarter
	if err != nil {
		return nil, err
	}
	svc.ApiUrl = parsedUrl

	return svc, nil
}

func (svc AccountsApiService) Create(accountData *models.AccountData) (*models.AccountData, error) {
	url := svc.ApiUrl
	err := validateCreate(url, accountData)
	if err != nil {
		return nil, err
	}

	buffer, err := encode(models.AccountDataRequest{Data: accountData})
	if err != nil {
		return nil, err
	}

	httpResponse, err := svc.HttpClient.Post(url.String(), "application/vnd.api+json", buffer)
	if err != nil {
		return nil, err
	}

	return svc.handleResponse(httpResponse, err)
}

// just placeholders - not in scope of this assigment
func validateCreate(*url.URL, *models.AccountData) error {
	return nil
}
func validateFetch(*url.URL) error {
	return nil
}
func validateDelete(*url.URL, *models.AccountData) error {
	return nil
}

func (svc AccountsApiService) handleResponse(httpResponse *http.Response, err error) (*models.AccountData, error) {
	if httpResponse.StatusCode >= http.StatusBadRequest {
		return nil, svc.handleError(err, httpResponse)
	}

	var accountDataResponse models.AccountDataResponse
	if httpResponse.Body != http.NoBody {
		err = json.NewDecoder(httpResponse.Body).Decode(&accountDataResponse)
		if err != nil {
			return nil, err
		}
	}
	return accountDataResponse.Data, nil
}

func encode(data any) (*bytes.Buffer, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(body), nil
}

func (svc AccountsApiService) Fetch(data *models.AccountData) (*models.AccountData, error) {
	url, err := url.Parse(svc.ApiUrl.String() + "/" + data.ID) // TODO: make it smarter
	if err != nil {
		return nil, err
	}

	err = validateFetch(url)
	if err != nil {
		return nil, err
	}

	httpResponse, err := svc.HttpClient.Get(url.String())
	if err != nil {
		return nil, err
	}

	return svc.handleResponse(httpResponse, err)
}

func (svc AccountsApiService) Delete(data *models.AccountData) error {
	url, err := url.Parse(svc.ApiUrl.String() + "/" + data.ID) // TODO: make it smarter
	if err != nil {
		return err
	}

	err = setParams(url, data)
	if err != nil {
		return err
	}

	err = validateDelete(url, data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodDelete, url.String(), bytes.NewReader(nil))
	if err != nil {
		return err
	}

	httpResponse, err := svc.HttpClient.Do(req)
	if err != nil {
		return err
	}

	_, err = svc.handleResponse(httpResponse, err)
	return err
}

func setParams(url *url.URL, data *models.AccountData) error {
	if data.Version == nil {
		return errors.New("data.Version cannot be nil")
	}
	setParam(url, "version", strconv.FormatInt(*data.Version, 10))
	return nil
}

func setParam(url *url.URL, key, value string) {
	values := url.Query()
	values.Set(key, value)
	url.RawQuery = values.Encode()
}

func (svc AccountsApiService) handleError(err error, httpResponse *http.Response) error {
	var errorResponse models.ErrorResponse
	err = json.NewDecoder(httpResponse.Body).Decode(&errorResponse)
	if err != nil {
		return err
	}

	return errors.New(errorResponse.ErrorMessage)
}
