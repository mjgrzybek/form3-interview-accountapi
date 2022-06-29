package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	utils "github.com/mjgrzybek/form3-interview-accountapi/client/internal/utils"
	"github.com/mjgrzybek/form3-interview-accountapi/client/pkg/models"
)

type accountsApiService struct {
	*Client
	Endpoint *url.URL
}

func newAccountsApiService(client *Client) *accountsApiService {
	svc := &accountsApiService{Client: client}
	svc.Endpoint = utils.JoinPathUrl(*svc.ApiUrl, "organisation", "accounts")
	return svc
}

func (svc accountsApiService) Create(ctx context.Context, accountData *models.AccountData) (*models.AccountData, error) {
	buffer, err := utils.Encode(models.AccountDataRequest{Data: accountData})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, svc.Endpoint.String(), buffer)
	if err != nil {
		return nil, err
	}
	req.Header.Set(HeaderNameContentType, HeaderValueVendorJson)

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

func (svc accountsApiService) Fetch(ctx context.Context, data *models.AccountIdVersion) (*models.AccountData, error) {
	url := utils.JoinPathUrl(*svc.Endpoint, data.ID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), bytes.NewReader(nil))
	if err != nil {
		return nil, err
	}
	req.Header.Set(HeaderNameAccept, HeaderValueVendorJson)

	err = validateFetch(req)
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

func (svc accountsApiService) Delete(ctx context.Context, data *models.AccountIdVersion) error {
	url := utils.JoinPathUrl(*svc.Endpoint, data.ID)

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

func (svc accountsApiService) handleResponse(httpResponse *http.Response) (*models.AccountData, error) {
	if httpResponse.StatusCode < http.StatusOK {
		return svc.handleInformation(httpResponse)
	}

	if httpResponse.StatusCode >= http.StatusOK && httpResponse.StatusCode < http.StatusMultipleChoices {
		return svc.handleSuccess(httpResponse)
	}

	if httpResponse.StatusCode >= http.StatusMultipleChoices && httpResponse.StatusCode < http.StatusBadRequest {
		return svc.handleRedirection(httpResponse)
	}

	if httpResponse.StatusCode >= http.StatusBadRequest && httpResponse.StatusCode < http.StatusInternalServerError {
		return svc.handleClientError(httpResponse)
	}
	if httpResponse.StatusCode >= http.StatusInternalServerError {
		return svc.handleServerError(httpResponse)
	}

	return nil, errors.New(fmt.Sprintf("Unable to interpret httpResponse: %+v", httpResponse))
}

func (svc accountsApiService) handleGeneric(httpResponse *http.Response) error {
	return errors.New(httpResponse.Status)
}

func (svc accountsApiService) handleInformation(httpResponse *http.Response) (*models.AccountData, error) {
	return nil, svc.handleGeneric(httpResponse)
}

func (svc accountsApiService) handleSuccess(httpResponse *http.Response) (*models.AccountData, error) {
	var accountDataResponse models.AccountDataResponse
	if httpResponse.Body != http.NoBody {
		err := json.NewDecoder(httpResponse.Body).Decode(&accountDataResponse)
		if err != nil {
			return nil, err
		}
	}
	return accountDataResponse.Data, nil
}
func (svc accountsApiService) handleRedirection(httpResponse *http.Response) (*models.AccountData, error) {
	return nil, svc.handleGeneric(httpResponse)
}

func (svc accountsApiService) handleClientError(httpResponse *http.Response) (*models.AccountData, error) {
	var errorResponse models.ErrorResponse
	err := json.NewDecoder(httpResponse.Body).Decode(&errorResponse)
	if err != nil {
		return nil, err
	}

	return nil, errors.New(errorResponse.ErrorMessage)
}

func (svc accountsApiService) handleServerError(httpResponse *http.Response) (*models.AccountData, error) {
	return nil, svc.handleGeneric(httpResponse)
}

func setParams(url *url.URL, data *models.AccountIdVersion) error {
	if data.Version == nil {
		return errors.New("version cannot be nil")
	}
	utils.SetParam(url, "version", strconv.FormatInt(*data.Version, 10))
	return nil
}
