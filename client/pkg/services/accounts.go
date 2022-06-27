package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	client "github.com/mjgrzybek/form3-interview-accountapi/client/internal/client"
	utils "github.com/mjgrzybek/form3-interview-accountapi/client/internal/utils"
	"github.com/mjgrzybek/form3-interview-accountapi/client/pkg/models"
)

type AccountsApiService client.Client

func NewAccountsApiService() (*AccountsApiService, error) {
	client, err := client.NewClient()
	if err != nil {
		return nil, err
	}
	svc := (*AccountsApiService)(client)
	svc.ApiUrl = utils.JoinPathUrl(*svc.ApiUrl, "organisation", "accounts")

	return svc, nil
}

func (svc AccountsApiService) Create(ctx context.Context, accountData *models.AccountData) (*models.AccountData, error) {
	buffer, err := utils.Encode(models.AccountDataRequest{Data: accountData})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, svc.ApiUrl.String(), buffer)
	if err != nil {
		return nil, err
	}
	req.Header.Set(client.HeaderNameContentType, client.HeaderValueVendorJson)

	err = validateCreate(req)
	if err != nil {
		return nil, err
	}

	httpResponse, err := svc.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	return svc.handleResponse(httpResponse)
}

func (svc AccountsApiService) Fetch(ctx context.Context, data *models.AccountData) (*models.AccountData, error) {
	url := utils.JoinPathUrl(*svc.ApiUrl, data.ID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), bytes.NewReader(nil))
	if err != nil {
		return nil, err
	}
	req.Header.Set(client.HeaderNameAccept, "application/vnd.api+json")

	validateFetch(req)

	httpResponse, err := svc.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	return svc.handleResponse(httpResponse)
}

func (svc AccountsApiService) Delete(ctx context.Context, data *models.AccountData) error {
	url := utils.JoinPathUrl(*svc.ApiUrl, data.ID)

	err := setParams(url, data)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url.String(), bytes.NewReader(nil))
	if err != nil {
		return err
	}

	err = validateDelete(req)
	if err != nil {
		return err
	}

	httpResponse, err := svc.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer httpResponse.Body.Close()

	_, err = svc.handleResponse(httpResponse)
	return err
}

func (svc AccountsApiService) handleResponse(httpResponse *http.Response) (*models.AccountData, error) {
	if httpResponse.StatusCode >= http.StatusBadRequest {
		return nil, svc.handleError(httpResponse)
	}

	var accountDataResponse models.AccountDataResponse
	if httpResponse.Body != http.NoBody {
		err := json.NewDecoder(httpResponse.Body).Decode(&accountDataResponse)
		if err != nil {
			return nil, err
		}
	}
	return accountDataResponse.Data, nil
}

func (svc AccountsApiService) handleError(httpResponse *http.Response) error {
	var errorResponse models.ErrorResponse
	err := json.NewDecoder(httpResponse.Body).Decode(&errorResponse)
	if err != nil {
		return err
	}

	return errors.New(errorResponse.ErrorMessage)
}

func setParams(url *url.URL, data *models.AccountData) error {
	if data.Version == nil {
		return errors.New("version cannot be nil")
	}
	utils.SetParam(url, "version", strconv.FormatInt(*data.Version, 10))
	return nil
}
