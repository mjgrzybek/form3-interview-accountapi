//go:build e2e

package services

import (
	"context"
	"errors"
	"log"
	"net/url"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/mjgrzybek/form3-interview-accountapi/client/internal/address"
	"github.com/mjgrzybek/form3-interview-accountapi/client/pkg/models"
	"github.com/stretchr/testify/assert"

	client "github.com/mjgrzybek/form3-interview-accountapi/client/pkg/services"
)

var apiUrl *url.URL

func init() {
	url, err := getApiUrlFromEnv()
	if err != nil {
		log.Fatalln(err)
	}
	apiUrl = url
}

func getApiUrlFromEnv() (*url.URL, error) {
	const ApiUrlEnvVarName = "API_URL"

	apiUrlEnvVarValue, ok := os.LookupEnv(ApiUrlEnvVarName)
	if !ok {
		return nil, errors.New("Environment variable \"" + ApiUrlEnvVarName + "\" not set")
	}

	apiUrl, err := url.Parse(apiUrlEnvVarValue)
	if err != nil {
		return nil, err
	}

	return apiUrl, nil
}

func TestAccountApi_Create(t *testing.T) {
	t.Run("create unique entry", func(t *testing.T) {
		svc, err := client.NewAccountsApiService(apiUrl)
		assert.NoError(t, err)

		accountRequestData := RequestsData["create"]
		accountResponseData, err := svc.Create(context.TODO(), accountRequestData)

		assert.Equal(t, ResponsesData["create"], accountResponseData)
		assert.NoError(t, err)
	})
	t.Run("create entry already existing", func(t *testing.T) {
		svc, err := client.NewAccountsApiService(apiUrl)
		assert.NoError(t, err)

		accountRequestData := RequestsData["create"]
		accountResponseData, err := svc.Create(context.TODO(), accountRequestData)

		assert.Nil(t, accountResponseData)
		assert.EqualError(t, err, "Account cannot be created as it violates a duplicate constraint")
	})
	t.Run("fail creating entry because server-side request validation failed", func(t *testing.T) {
		svc, err := client.NewAccountsApiService(apiUrl)
		assert.NoError(t, err)

		accountRequestData := RequestsData["create"]
		accountRequestData.Attributes.Country = address.Of("This country doesn't exist")
		accountResponseData, err := svc.Create(context.TODO(), accountRequestData)

		assert.Nil(t, accountResponseData)
		assert.EqualError(t, err, "validation failure list:\nvalidation failure list:\nvalidation failure list:\ncountry in body should match '^[A-Z]{2}$'")
	})
}

func TestAccountApi_Fetch(t *testing.T) {
	t.Run("fetch existing entry", func(t *testing.T) {
		svc, err := client.NewAccountsApiService(apiUrl)
		assert.NoError(t, err)

		accountRequestData := RequestsData["fetch"]
		accountResponseData, err := svc.Fetch(context.TODO(), accountRequestData)
		assert.Equal(t, ResponsesData["fetch"], accountResponseData)
		assert.NoError(t, err)
	})
	t.Run("fail fetching non-existing entry", func(t *testing.T) {
		svc, err := client.NewAccountsApiService(apiUrl)
		assert.NoError(t, err)

		accountRequestData := RequestsData["fetch"]
		setRandomUuid(t, accountRequestData)

		accountResponseData, err := svc.Fetch(context.TODO(), accountRequestData)

		assert.Nil(t, accountResponseData)
		assert.EqualError(t, err, "record "+accountRequestData.ID+" does not exist")
	})
}

func TestAccountApi_Delete(t *testing.T) {
	t.Run("delete existing account", func(t *testing.T) {
		svc, err := client.NewAccountsApiService(apiUrl)
		assert.NoError(t, err)

		accountRequestData := RequestsData["delete"]
		err = svc.Delete(context.TODO(), accountRequestData)
		assert.NoError(t, err)
	})
	t.Run("fail deleting non-existing account", func(t *testing.T) {
		svc, err := client.NewAccountsApiService(apiUrl)
		assert.NoError(t, err)

		accountRequestData := RequestsData["delete"]
		setRandomUuid(t, accountRequestData)

		err = svc.Delete(context.TODO(), accountRequestData)
		assert.EqualError(t, err, "EOF")
	})
}

func setRandomUuid(t *testing.T, accountRequestData *models.AccountData) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		t.Skipf("Unable to create new UUID %v", err)
	}
	accountRequestData.ID = uuid.String()
}
