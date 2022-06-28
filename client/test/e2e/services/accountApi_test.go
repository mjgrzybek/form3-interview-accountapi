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
	"github.com/mjgrzybek/form3-interview-accountapi/client/pkg/client"
	"github.com/mjgrzybek/form3-interview-accountapi/client/pkg/models"
	"github.com/stretchr/testify/assert"
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
		client := client.NewClient(apiUrl)

		accountResponseData, err := client.AccountApi.Create(context.TODO(), CreateRequestsData)

		assert.Equal(t, CreateResponseData, accountResponseData)
		assert.NoError(t, err)
	})
	t.Run("create entry already existing", func(t *testing.T) {
		client := client.NewClient(apiUrl)

		accountResponseData, err := client.AccountApi.Create(context.TODO(), CreateResponseData)

		assert.Nil(t, accountResponseData)
		assert.EqualError(t, err, "Account cannot be created as it violates a duplicate constraint")
	})
	t.Run("fail creating entry because server-side request validation failed", func(t *testing.T) {
		client := client.NewClient(apiUrl)

		CreateRequestsData.Attributes.Country = address.Of("This country doesn't exist")
		accountResponseData, err := client.AccountApi.Create(context.TODO(), CreateRequestsData)

		assert.Nil(t, accountResponseData)
		assert.EqualError(t, err, "validation failure list:\nvalidation failure list:\nvalidation failure list:\ncountry in body should match '^[A-Z]{2}$'")
	})
}

func TestAccountApi_Fetch(t *testing.T) {
	t.Run("fetch existing entry", func(t *testing.T) {
		client := client.NewClient(apiUrl)

		accountResponseData, err := client.AccountApi.Fetch(context.TODO(), FetchRequestsData)
		assert.Equal(t, FetchResponseData, accountResponseData)
		assert.NoError(t, err)
	})
	t.Run("fail fetching non-existing entry", func(t *testing.T) {
		client := client.NewClient(apiUrl)

		setRandomUuid(t, FetchRequestsData)

		accountResponseData, err := client.AccountApi.Fetch(context.TODO(), FetchRequestsData)

		assert.Nil(t, accountResponseData)
		assert.EqualError(t, err, "record "+FetchRequestsData.ID+" does not exist")
	})
}

func TestAccountApi_Delete(t *testing.T) {
	t.Run("delete existing account", func(t *testing.T) {
		client := client.NewClient(apiUrl)

		err := client.AccountApi.Delete(context.TODO(), DeleteRequestData)

		assert.NoError(t, err)
	})
	t.Run("fail deleting non-existing account", func(t *testing.T) {
		client := client.NewClient(apiUrl)

		setRandomUuid(t, DeleteRequestData)

		err := client.AccountApi.Delete(context.TODO(), DeleteRequestData)
		assert.EqualError(t, err, "EOF")
	})
}

func setRandomUuid(t *testing.T, accountRequestData *models.AccountIdVersion) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		t.Skipf("Unable to create new UUID %v", err)
	}
	accountRequestData.ID = uuid.String()
}
